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
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unichainplatform/unichain/blockchain"
	"github.com/unichainplatform/unichain/cmd/utils"
	"github.com/unichainplatform/unichain/debug"
	"github.com/unichainplatform/unichain/uniservice"
	"github.com/unichainplatform/unichain/metrics"
	"github.com/unichainplatform/unichain/metrics/influxdb"
	"github.com/unichainplatform/unichain/node"
)

var (
	errNoConfigFile    string
	errViperReadConfig error
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "uni",
	Short: "uni is a Leading High-performance Ledger",
	Long:  `uni is a Leading High-performance Ledger`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if viper.ConfigFileUsed() != "" {
			err = viper.Unmarshal(uniCfgInstance)
		}
		uniCfgInstance.LogCfg.Setup()
		if errNoConfigFile != "" {
			log.Info(errNoConfigFile)
		}

		if errViperReadConfig != nil {
			log.Error("Can't read config, use default configuration", "err", errViperReadConfig)
		}

		if err != nil {
			log.Error("viper umarshal config file faild", "err", err)
		}

		if err := debug.Setup(uniCfgInstance.DebugCfg); err != nil {
			log.Error("debug setup faild", "err", err)
		}

		log.Info("unichain node", "version", utils.FullVersion())

		node, err := makeNode()
		if err != nil {
			log.Error("uni make node failed.", "err", err)
			return
		}

		if err := registerService(node); err != nil {
			log.Error("uni register service failed.", "err", err)
			return
		}

		if err := startNode(node); err != nil {
			log.Error("uni start node failed.", "err", err)
			return
		}

		node.Wait()
		debug.Exit()
	},
}

func makeNode() (*node.Node, error) {
	genesis := blockchain.DefaultGenesis()
	// set miner config
	SetupMetrics()
	// Make sure we have a valid genesis JSON
	if len(uniCfgInstance.GenesisFile) != 0 {
		log.Info("Reading read genesis file", "path", uniCfgInstance.GenesisFile)
		file, err := os.Open(uniCfgInstance.GenesisFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to read genesis file: %v(%v)", uniCfgInstance.GenesisFile, err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(genesis); err != nil {
			return nil, fmt.Errorf("invalid genesis file: %v(%v)", uniCfgInstance.GenesisFile, err)
		}
		uniCfgInstance.UniServiceCfg.Genesis = genesis

	}
	block, _, err := genesis.ToBlock(nil)
	if err != nil {
		return nil, err
	}
	// p2p used to generate MagicNetID
	uniCfgInstance.NodeCfg.P2PConfig.GenesisHash = block.Hash()
	return node.New(uniCfgInstance.NodeCfg)
}

// SetupMetrics set metrics
func SetupMetrics() {
	//need to set metrice.Enabled = true in metrics source code
	if uniCfgInstance.UniServiceCfg.MetricsConf.MetricsFlag {
		log.Info("Enabling metrics collection")
		if uniCfgInstance.UniServiceCfg.MetricsConf.InfluxDBFlag {
			log.Info("Enabling influxdb collection")
			go influxdb.InfluxDBWithTags(metrics.DefaultRegistry, 10*time.Second, uniCfgInstance.UniServiceCfg.MetricsConf.URL,
				uniCfgInstance.UniServiceCfg.MetricsConf.DataBase, uniCfgInstance.UniServiceCfg.MetricsConf.UserName, uniCfgInstance.UniServiceCfg.MetricsConf.PassWd,
				uniCfgInstance.UniServiceCfg.MetricsConf.NameSpace, map[string]string{})
		}

	}
}

// start up the node itself
func startNode(stack *node.Node) error {
	debug.Memsize.Add("node", stack)
	if err := stack.Start(); err != nil {
		return err
	}
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		log.Info("Got interrupt, shutting down...")
		go stack.Stop()
		for i := 10; i > 0; i-- {
			<-sigc
			if i > 1 {
				log.Warn("Already shutting down, interrupt more to panic.", "times", i-1)
			}
		}
		debug.Exit() // ensure trace and CPU profile data is flushed.
		debug.LoudPanic("boom")
	}()
	return nil
}

func registerService(stack *node.Node) error {
	return stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return uniservice.New(ctx, uniCfgInstance.UniServiceCfg)
	})
}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		errNoConfigFile = "No config file , use default configuration."
		return
	}
	if err := viper.ReadInConfig(); err != nil {
		errViperReadConfig = err
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(utils.VersionCmd)
	addFlags(RootCmd.Flags())
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
