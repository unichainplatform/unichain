// Copyright 2018 The UniChain Team Authors
// This file is part of the unichain project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package processor

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/unichainplatform/unichain/accountmanager"
	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/consensus"
	"github.com/unichainplatform/unichain/params"
	"github.com/unichainplatform/unichain/processor/vm"
	"github.com/unichainplatform/unichain/state"
	"github.com/unichainplatform/unichain/types"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	bc     ChainContext      // Canonical block chain
	engine consensus.IEngine // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(bc ChainContext, engine consensus.IEngine) *StateProcessor {
	return &StateProcessor{
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) ([]*types.Receipt, []*types.Log, uint64, error) {
	var (
		receipts []*types.Receipt
		usedGas  = new(uint64)
		header   = block.Header()
		allLogs  []*types.Log
		gp       = new(common.GasPool).AddGas(block.GasLimit())
	)

	// Prepare the block, applying any consensus engine specific extras (e.g. update last)
	p.engine.Prepare(p.bc, header, block.Transactions(), receipts, statedb)

	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		receipt, _, err := p.ApplyTransaction(nil, gp, statedb, header, tx, usedGas, cfg)
		if err != nil {
			return nil, nil, 0, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, block.Transactions(), receipts, statedb)

	return receipts, allLogs, *usedGas, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func (p *StateProcessor) ApplyTransaction(author *common.Name, gp *common.GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) (*types.Receipt, uint64, error) {
	bc := p.bc
	config := bc.Config()
	accountDB, err := accountmanager.NewAccountManager(statedb)
	if err != nil {
		return nil, 0, err
	}

	// todo for the moment，only system asset
	// assetID := tx.GasAssetID()
	assetID := p.bc.Config().SysTokenID
	if assetID != tx.GasAssetID() {
		return nil, 0, fmt.Errorf("only support system asset %d as tx fee", p.bc.Config().SysTokenID)
	}
	//timer for vm exec overtime
	var t *time.Timer
	//
	var totalGas uint64
	var ios []*types.ActionResult
	detailTx := &types.DetailTx{}
	var detailActions []*types.DetailAction
	for i, action := range tx.GetActions() {
		if needCheckSign(accountDB, action) {
			if err := accountDB.RecoverTx(types.NewSigner(config.ChainID), tx); err != nil {
				return nil, 0, err
			}
		}

		nonce, err := accountDB.GetNonce(action.Sender())
		if err != nil {
			return nil, 0, err
		}
		if nonce < action.Nonce() {
			return nil, 0, ErrNonceTooHigh
		} else if nonce > action.Nonce() {
			return nil, 0, ErrNonceTooLow
		}

		var gasPayer = action.Sender()
		var gasPrice = tx.GasPrice()
		if tx.PayerExist() {
			if header.CurForkID() >= params.ForkID4 {
				gasPayer = action.Payer()
				gasPrice = action.PayerGasPrice()
			} else {
				return nil, 0, errPayerNotSupport
			}
		} else {
			if action.PayerIsExist() {
				return nil, 0, errPayerNotSupport
			}
		}

		evmcontext := &EvmContext{
			ChainContext:  p.bc,
			EngineContext: p.engine,
		}
		context := NewEVMContext(action.Sender(), action.Recipient(), assetID, gasPrice, header, evmcontext, author)
		vmenv := vm.NewEVM(context, accountDB, statedb, config, cfg)

		//will abort the vm if overtime
		if false == cfg.EndTime.IsZero() {
			t = time.AfterFunc(cfg.EndTime.Sub(time.Now()), func() {
				vmenv.OverTimeAbort()
			})
		}

		_, gas, failed, err, vmerr := ApplyMessage(accountDB, vmenv, action, gp, gasPrice, gasPayer, assetID, config, p.engine)

		if false == cfg.EndTime.IsZero() {
			//close timer
			t.Stop()
		}

		if err != nil {
			return nil, 0, err
		}

		*usedGas += gas
		totalGas += gas

		var status uint64
		if failed {
			status = types.ReceiptStatusFailed
		} else {
			status = types.ReceiptStatusSuccessful

		}
		vmerrstr := ""
		if vmerr != nil {
			vmerrstr = vmerr.Error()
			log.Debug("processer apply transaction ", "hash", tx.Hash(), "err", vmerrstr)
		}
		var gasAllot []*types.GasDistribution
		for key, gas := range vmenv.FounderGasMap {
			gasAllot = append(gasAllot, &types.GasDistribution{Account: key.ObjectName.String(), Gas: uint64(gas.Value), TypeID: gas.TypeID})
		}
		ios = append(ios, &types.ActionResult{Status: status, Index: uint64(i), GasUsed: gas, GasAllot: gasAllot, Error: vmerrstr})

		internalTxLog := make([]*types.InternalAction, 0, len(vmenv.InternalTxs))
		for _, internalAction := range vmenv.InternalTxs {
			internalAction.Action.SetHash(action.Hash())
			internalTxLog = append(internalTxLog, internalAction)
		}
		detailActions = append(detailActions, &types.DetailAction{InternalActions: internalTxLog})
	}
	root := statedb.ReceiptRoot()
	receipt := types.NewReceipt(root[:], *usedGas, totalGas)
	receipt.TxHash = tx.Hash()
	receipt.ActionResults = ios
	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom([]*types.Receipt{receipt})

	detailTx.TxHash = receipt.TxHash
	detailTx.Actions = detailActions
	receipt.SetInternalTxsLog(detailTx)
	return receipt, totalGas, nil
}

func needCheckSign(accountDB *accountmanager.AccountManager, action *types.Action) bool {
	authorVersion := types.GetAuthorCache(action)
	if len(authorVersion) == 0 {
		return true
	}
	for name, version := range authorVersion {
		if tmpVersion, err := accountDB.GetAuthorVersion(name); err != nil || version != tmpVersion {
			return true
		}
	}
	return false
}
