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

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unichainplatform/unichain/blockchain"
	"github.com/unichainplatform/unichain/uniservice"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init -g <genesis> -d <datadir>",
	Short: "Bootstrap and initialize a new genesis block",
	Long:  `Bootstrap and initialize a new genesis block`,
	Run: func(cmd *cobra.Command, args []string) {
		uniCfgInstance.LogCfg.Setup()
		if err := initGenesis(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&uniCfgInstance.GenesisFile, "genesis", "g", "", "Genesis json file")
	initCmd.Flags().StringVarP(&uniCfgInstance.NodeCfg.DataDir, "datadir", "d", uniCfgInstance.NodeCfg.DataDir, "Data directory for the databases ")
}

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func initGenesis() error {
	// Make sure we have a valid genesis JSON
	genesis := blockchain.DefaultGenesis()
	if len(uniCfgInstance.GenesisFile) != 0 {
		file, err := os.Open(uniCfgInstance.GenesisFile)
		if err != nil {
			return fmt.Errorf("Failed to read genesis file: %v(%v)", uniCfgInstance.GenesisFile, err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(genesis); err != nil {
			return fmt.Errorf("invalid genesis file: %v(%v)", uniCfgInstance.GenesisFile, err)
		}
	}

	stack, err := makeNode()
	if err != nil {
		return err
	}

	_, err = uniservice.New(stack.GetNodeConfig(), uniCfgInstance.UniServiceCfg)
	if err != nil {
		return err
	}
	return nil
}
