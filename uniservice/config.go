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
	"github.com/unichainplatform/unichain/blockchain"
	"github.com/unichainplatform/unichain/uniservice/gasprice"
	"github.com/unichainplatform/unichain/metrics"
	"github.com/unichainplatform/unichain/txpool"
)

// Config uniservice config
type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the main net block is used.
	Genesis *blockchain.Genesis `toml:",omitempty"`

	// Database options
	DatabaseHandles int
	DatabaseCache   int `mapstructure:"databasecache"`

	// Transaction pool options
	TxPool *txpool.Config `mapstructure:"txpool"`

	// Gas Price Oracle options
	GasPrice gasprice.Config `mapstructure:"gpo"`

	// miner
	Miner *MinerConfig `mapstructure:"miner"`

	MetricsConf *metrics.Config `mapstructure:"metrics"`

	StatePruning    bool `mapstructure:"statepruning"`
	ContractLogFlag bool `mapstructure:"contractlog"`

	BadHashes   []string `mapstructure:"badhashes"`
	StartNumber uint64   `mapstructure:"startnumber"`
}

// MinerConfig miner config
type MinerConfig struct {
	Start       bool     `mapstructure:"start"`
	Delay       uint64   `mapstructure:"delay"`
	Name        string   `mapstructure:"name"`
	PrivateKeys []string `mapstructure:"private"`
	ExtraData   string   `mapstructure:"extra"`
}
