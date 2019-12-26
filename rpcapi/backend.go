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

// package rpcapi implements the general API functions.

package rpcapi

import (
	"context"
	"math/big"

	"github.com/unichainplatform/unichain/accountmanager"
	"github.com/unichainplatform/unichain/blockchain"
	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/consensus"
	"github.com/unichainplatform/unichain/debug"
	"github.com/unichainplatform/unichain/feemanager"
	"github.com/unichainplatform/unichain/params"
	"github.com/unichainplatform/unichain/processor/vm"
	"github.com/unichainplatform/unichain/rpc"
	"github.com/unichainplatform/unichain/rpcapi/filters"
	"github.com/unichainplatform/unichain/state"
	"github.com/unichainplatform/unichain/txpool"
	"github.com/unichainplatform/unichain/types"
	"github.com/unichainplatform/unichain/utils/fdb"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// uniservice API
	ChainDb() fdb.Database
	ChainConfig() *params.ChainConfig
	SuggestPrice(ctx context.Context) (*big.Int, error)

	// BlockChain API
	CurrentBlock() *types.Block
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) *types.Header
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) *types.Block
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) ([]*types.Receipt, error)
	GetDetailTxsLog(ctx context.Context, hash common.Hash) ([]*types.DetailTx, error)
	GetBlockDetailLog(ctx context.Context, blockNr rpc.BlockNumber) *types.BlockAndResult
	GetTd(blockHash common.Hash) *big.Int
	GetEVM(ctx context.Context, account *accountmanager.AccountManager, state *state.StateDB, from common.Name, to common.Name, assetID uint64, gasPrice *big.Int, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error)
	GetDetailTxByFilter(ctx context.Context, filterFn func(common.Name) bool, blockNr, lookbackNum uint64) []*types.DetailTx
	GetTxsByFilter(ctx context.Context, filterFn func(common.Name) bool, blockNr, lookbackNum uint64) *types.AccountTxs
	GetBadBlocks(ctx context.Context) ([]*types.Block, error)
	ForkStatus(statedb *state.StateDB) (*blockchain.ForkConfig, blockchain.ForkInfo, error)
	SetStatePruning(enable bool) (bool, uint64)

	// TxPool
	TxPool() *txpool.TxPool
	SendTx(ctx context.Context, signedTx *types.Transaction) error

	SetGasPrice(gasPrice *big.Int) bool

	//Account API
	GetAccountManager() (*accountmanager.AccountManager, error)

	//fee manager
	GetFeeManager() (*feemanager.FeeManager, error)
	GetFeeManagerByTime(time uint64) (*feemanager.FeeManager, error)

	// P2P
	AddPeer(url string) error
	RemovePeer(url string) error
	AddTrustedPeer(url string) error
	RemoveTrustedPeer(url string) error
	SeedNodes() []string
	PeerCount() int
	Peers() []string
	BadNodesCount() int
	BadNodes() []string
	AddBadNode(url string) error
	RemoveBadNode(url string) error
	SelfNode() string
	Engine() consensus.IEngine
	APIs() []rpc.API

	// Filter Log
	HeaderByHash(ctx context.Context, blockHash common.Hash) *types.Header
	GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error)
}

func GetAPIs(apiBackend Backend) []rpc.API {
	apis := []rpc.API{
		{
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPrivateTxPoolAPI(apiBackend),
		},
		{
			Namespace: "bc",
			Version:   "1.0",
			Service:   NewPrivateBlockChainAPI(apiBackend),
		},
		{
			Namespace: "uni",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "uni",
			Version:   "1.0",
			Service:   NewPublicUniChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "uni",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(apiBackend),
			Public:    true,
		},
		{
			Namespace: "account",
			Version:   "1.0",
			Service:   NewAccountAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "fee",
			Version:   "1.0",
			Service:   NewFeeAPI(apiBackend),
			Public:    true,
		},
		{
			Namespace: "p2p",
			Version:   "1.0",
			Service:   NewPrivateP2pAPI(apiBackend),
		},
		{
			Namespace: "debug",
			Version:   "1.0",
			Service:   debug.Handler,
		},
	}
	return append(apis, apiBackend.APIs()...)
}
