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

package uniservice

import (
	"math/big"

	"github.com/ethereum/go-ethereum/log"
	"github.com/unichainplatform/unichain/blockchain"
	"github.com/unichainplatform/unichain/consensus"
	"github.com/unichainplatform/unichain/consensus/dpos"
	"github.com/unichainplatform/unichain/consensus/miner"
	"github.com/unichainplatform/unichain/uniservice/gasprice"
	"github.com/unichainplatform/unichain/node"
	"github.com/unichainplatform/unichain/p2p"
	adaptor "github.com/unichainplatform/unichain/p2p/protoadaptor"
	"github.com/unichainplatform/unichain/params"
	"github.com/unichainplatform/unichain/processor"
	"github.com/unichainplatform/unichain/processor/vm"
	"github.com/unichainplatform/unichain/rpc"
	"github.com/unichainplatform/unichain/rpcapi"
	"github.com/unichainplatform/unichain/txpool"
	"github.com/unichainplatform/unichain/utils/fdb"
)

// UniService implements the unichain service.
type UniService struct {
	config       *Config
	chainConfig  *params.ChainConfig
	shutdownChan chan bool // Channel for shutting down the service
	blockchain   *blockchain.BlockChain
	txPool       *txpool.TxPool
	chainDb      fdb.Database // Block chain database
	engine       consensus.IEngine
	miner        *miner.Miner
	p2pServer    *adaptor.ProtoAdaptor
	APIBackend   *APIBackend
}

// New creates a new uniservice object (including the initialisation of the common uniservice object)
func New(ctx *node.ServiceContext, config *Config) (*UniService, error) {
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}

	chainCfg, dposCfg, _, err := blockchain.SetupGenesisBlock(chainDb, config.Genesis)
	if err != nil {
		return nil, err
	}

	ctx.AppendBootNodes(chainCfg.BootNodes)

	uniService := &UniService{
		config:       config,
		chainDb:      chainDb,
		chainConfig:  chainCfg,
		p2pServer:    ctx.P2P,
		shutdownChan: make(chan bool),
	}

	//blockchain
	vmconfig := vm.Config{
		ContractLogFlag: config.ContractLogFlag,
	}

	uniService.blockchain, err = blockchain.NewBlockChain(chainDb, config.StatePruning, vmconfig, uniService.chainConfig, config.BadHashes, config.StartNumber, txpool.SenderCacher)
	if err != nil {
		return nil, err
	}

	// txpool
	if config.TxPool.Journal != "" {
		config.TxPool.Journal = ctx.ResolvePath(config.TxPool.Journal)
	}

	uniService.txPool = txpool.New(*config.TxPool, uniService.chainConfig, uniService.blockchain)

	engine := dpos.New(dposCfg, uniService.blockchain)
	uniService.engine = engine

	type bc struct {
		*blockchain.BlockChain
		consensus.IEngine
		*txpool.TxPool
		processor.Processor
	}

	bcc := &bc{
		uniService.blockchain,
		uniService.engine,
		uniService.txPool,
		nil,
	}

	validator := processor.NewBlockValidator(bcc, uniService.engine)
	txProcessor := processor.NewStateProcessor(bcc, uniService.engine)

	uniService.blockchain.SetValidator(validator)
	uniService.blockchain.SetProcessor(txProcessor)

	bcc.Processor = txProcessor
	uniService.miner = miner.NewMiner(bcc)
	uniService.miner.SetDelayDuration(config.Miner.Delay)
	uniService.miner.SetCoinbase(config.Miner.Name, config.Miner.PrivateKeys)
	uniService.miner.SetExtra([]byte(config.Miner.ExtraData))
	if config.Miner.Start {
		uniService.miner.Start(false)
	}

	uniService.APIBackend = &APIBackend{uniService: uniService}

	uniService.SetGasPrice(uniService.TxPool().GasPrice())
	return uniService, nil
}

// APIs return the collection of RPC services the uniservice package offers.
func (fs *UniService) APIs() []rpc.API {
	return rpcapi.GetAPIs(fs.APIBackend)
}

// Start implements node.Service, starting all internal goroutines.
func (fs *UniService) Start() error {
	log.Info("start unichain service...")
	return nil
}

// Stop implements node.Service, terminating all internal goroutine
func (fs *UniService) Stop() error {
	fs.miner.Stop()
	fs.blockchain.Stop()
	fs.txPool.Stop()
	fs.chainDb.Close()
	close(fs.shutdownChan)
	log.Info("uniservice stopped")
	return nil
}

func (fs *UniService) GasPrice() *big.Int {
	return fs.txPool.GasPrice()
}

func (fs *UniService) SetGasPrice(gasPrice *big.Int) bool {
	fs.config.GasPrice.Default = new(big.Int).SetBytes(gasPrice.Bytes())
	fs.APIBackend.gpo = gasprice.NewOracle(fs.APIBackend, fs.config.GasPrice)
	fs.txPool.SetGasPrice(new(big.Int).SetBytes(gasPrice.Bytes()))
	return true
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (fdb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *UniService) BlockChain() *blockchain.BlockChain { return s.blockchain }
func (s *UniService) TxPool() *txpool.TxPool             { return s.txPool }
func (s *UniService) Engine() consensus.IEngine          { return s.engine }
func (s *UniService) ChainDb() fdb.Database              { return s.chainDb }
func (s *UniService) Protocols() []p2p.Protocol          { return nil }
