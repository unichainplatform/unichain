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
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// ConfigFile the UniChain config file
	ConfigFile string
)

func addFlags(flags *flag.FlagSet) {
	// debug
	flags.BoolVar(
		&uniCfgInstance.DebugCfg.Pprof,
		"debug_pprof",
		uniCfgInstance.DebugCfg.Pprof,
		"Enable the pprof HTTP server",
	)
	viper.BindPFlag("debug.pprof", flags.Lookup("debug_pprof"))

	flags.IntVar(
		&uniCfgInstance.DebugCfg.PprofPort,
		"debug_pprof_port",
		uniCfgInstance.DebugCfg.PprofPort,
		"Pprof HTTP server listening port",
	)
	viper.BindPFlag("debug.pprofport", flags.Lookup("debug_pprof_port"))

	flags.StringVar(
		&uniCfgInstance.DebugCfg.PprofAddr,
		"debug_pprof_addr",
		uniCfgInstance.DebugCfg.PprofAddr,
		"Pprof HTTP server listening interface",
	)
	viper.BindPFlag("debug.pprofaddr", flags.Lookup("debug_pprof_addr"))

	flags.IntVar(
		&uniCfgInstance.DebugCfg.Memprofilerate,
		"debug_memprofilerate",
		uniCfgInstance.DebugCfg.Memprofilerate,
		"Turn on memory profiling with the given rate",
	)
	viper.BindPFlag("debug.memprofilerate", flags.Lookup("debug_memprofilerate"))

	flags.IntVar(
		&uniCfgInstance.DebugCfg.Blockprofilerate,
		"debug_blockprofilerate",
		uniCfgInstance.DebugCfg.Blockprofilerate,
		"Turn on block profiling with the given rate",
	)
	viper.BindPFlag("debug.blockprofilerate", flags.Lookup("debug_blockprofilerate"))

	flags.StringVar(
		&uniCfgInstance.DebugCfg.Cpuprofile,
		"debug_cpuprofile",
		uniCfgInstance.DebugCfg.Cpuprofile,
		"Write CPU profile to the given file",
	)
	viper.BindPFlag("debug.cpuprofile", flags.Lookup("debug_cpuprofile"))

	flags.StringVar(
		&uniCfgInstance.DebugCfg.Trace,
		"debug_trace",
		uniCfgInstance.DebugCfg.Trace,
		"Write execution trace to the given file",
	)
	viper.BindPFlag("debug.trace", flags.Lookup("debug_trace"))

	// log
	flags.StringVar(
		&uniCfgInstance.LogCfg.Logdir,
		"log_dir",
		uniCfgInstance.LogCfg.Logdir,
		"Writes log records to file chunks at the given path",
	)
	viper.BindPFlag("log.dir", flags.Lookup("log_dir"))

	flags.BoolVar(
		&uniCfgInstance.LogCfg.PrintOrigins,
		"log_debug",
		uniCfgInstance.LogCfg.PrintOrigins,
		"Prepends log messages with call-site location (file and line number)",
	)
	viper.BindPFlag("log.debug", flags.Lookup("log_debug"))

	flags.IntVar(
		&uniCfgInstance.LogCfg.Level,
		"log_level",
		uniCfgInstance.LogCfg.Level,
		"Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
	)
	viper.BindPFlag("log.level", flags.Lookup("log_level"))

	flags.StringVar(
		&uniCfgInstance.LogCfg.Vmodule,
		"log_module",
		uniCfgInstance.LogCfg.Vmodule,
		"Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. uni/*=5,p2p=4)",
	)
	viper.BindPFlag("log.module", flags.Lookup("log_module"))

	flags.StringVar(
		&uniCfgInstance.LogCfg.BacktraceAt,
		"log_backtrace",
		uniCfgInstance.LogCfg.BacktraceAt,
		"Request a stack trace at a specific logging statement (e.g. \"block.go:271\")",
	)
	viper.BindPFlag("log.backtrace", flags.Lookup("log_backtrace"))

	// config file
	flags.StringVarP(
		&ConfigFile,
		"config", "c",
		"",
		"TOML/YAML configuration file",
	)

	// Genesis File
	flags.StringVarP(
		&uniCfgInstance.GenesisFile,
		"genesis",
		"g", "",
		"Genesis json file",
	)
	viper.BindPFlag("genesis", flags.Lookup("genesis"))

	// node datadir
	flags.StringVarP(
		&uniCfgInstance.NodeCfg.DataDir,
		"datadir", "d",
		uniCfgInstance.NodeCfg.DataDir,
		"Data directory for the databases ",
	)
	viper.BindPFlag("node.datadir", flags.Lookup("datadir"))

	// node
	flags.StringVar(
		&uniCfgInstance.NodeCfg.IPCPath,
		"ipcpath",
		uniCfgInstance.NodeCfg.IPCPath,
		"RPC:ipc file name",
	)
	viper.BindPFlag("node.ipcpath", flags.Lookup("ipcpath"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.HTTPHost,
		"http_host",
		uniCfgInstance.NodeCfg.HTTPHost,
		"RPC:http host address",
	)
	viper.BindPFlag("node.httphost", flags.Lookup("http_host"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.HTTPPort,
		"http_port",
		uniCfgInstance.NodeCfg.HTTPPort,
		"RPC:http host port",
	)
	viper.BindPFlag("node.httpport", flags.Lookup("http_port"))

	flags.StringSliceVar(
		&uniCfgInstance.NodeCfg.HTTPModules,
		"http_modules",
		uniCfgInstance.NodeCfg.HTTPModules,
		"RPC:http api's offered over the HTTP-RPC interface",
	)
	viper.BindPFlag("node.httpmodules", flags.Lookup("http_modules"))

	flags.StringSliceVar(
		&uniCfgInstance.NodeCfg.HTTPCors,
		"http_cors",
		uniCfgInstance.NodeCfg.HTTPCors,
		"RPC:Which to accept cross origin",
	)
	viper.BindPFlag("node.httpcors", flags.Lookup("http_cors"))

	flags.StringSliceVar(
		&uniCfgInstance.NodeCfg.HTTPVirtualHosts,
		"http_vhosts",
		uniCfgInstance.NodeCfg.HTTPVirtualHosts,
		"RPC:http virtual hostnames from which to accept requests",
	)
	viper.BindPFlag("node.httpvirtualhosts", flags.Lookup("http_vhosts"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.WSHost,
		"ws_host",
		uniCfgInstance.NodeCfg.WSHost,
		"RPC:websocket host address",
	)
	viper.BindPFlag("node.wshost", flags.Lookup("ws_host"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.WSPort,
		"ws_port",
		uniCfgInstance.NodeCfg.WSPort,
		"RPC:websocket host port",
	)
	viper.BindPFlag("node.wsport", flags.Lookup("ws_port"))

	flags.StringSliceVar(
		&uniCfgInstance.NodeCfg.WSModules,
		"ws_modules",
		uniCfgInstance.NodeCfg.WSModules,
		"RPC:ws api's offered over the WS-RPC interface",
	)
	viper.BindPFlag("node.wsmodules", flags.Lookup("ws_modules"))

	flags.StringSliceVar(
		&uniCfgInstance.NodeCfg.WSOrigins,
		"ws_origins",
		uniCfgInstance.NodeCfg.WSOrigins,
		"RPC:ws origins from which to accept websockets requests",
	)
	viper.BindPFlag("node.wsorigins", flags.Lookup("ws_origins"))

	flags.BoolVar(
		&uniCfgInstance.NodeCfg.WSExposeAll,
		"ws_exposeall",
		uniCfgInstance.NodeCfg.WSExposeAll,
		"RPC:ws exposes all API modules via the WebSocket RPC interface rather than just the public ones.",
	)
	viper.BindPFlag("node.wsexposeall", flags.Lookup("ws_exposeall"))

	// uniservice database options
	flags.IntVar(
		&uniCfgInstance.UniServiceCfg.DatabaseCache,
		"database_cache",
		uniCfgInstance.UniServiceCfg.DatabaseCache,
		"Megabytes of memory allocated to internal database caching",
	)
	viper.BindPFlag("uniservice.databasecache", flags.Lookup("database_cache"))

	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.ContractLogFlag,
		"contractlog",
		uniCfgInstance.UniServiceCfg.ContractLogFlag,
		"flag for db to store contrat internal transaction log.",
	)
	viper.BindPFlag("uniservice.contractlog", flags.Lookup("contractlog"))

	// state pruning
	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.StatePruning,
		"statepruning_enable",
		uniCfgInstance.UniServiceCfg.StatePruning,
		"flag for enable/disable state pruning.",
	)
	viper.BindPFlag("uniservice.statepruning", flags.Lookup("statepruning_enable"))

	// start number
	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.StartNumber,
		"start_number",
		uniCfgInstance.UniServiceCfg.StartNumber,
		"start chain with a specified block number.",
	)
	viper.BindPFlag("uniservice.startnumber", flags.Lookup("start_number"))

	// add bad block hashs
	flags.StringSliceVar(
		&uniCfgInstance.UniServiceCfg.BadHashes,
		"bad_hashes",
		uniCfgInstance.UniServiceCfg.BadHashes,
		"blockchain refuse bad block hashes",
	)
	viper.BindPFlag("uniservice.badhashes", flags.Lookup("bad_hashes"))

	// txpool
	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.TxPool.NoLocals,
		"txpool_nolocals",
		uniCfgInstance.UniServiceCfg.TxPool.NoLocals,
		"Disables price exemptions for locally submitted transactions",
	)
	viper.BindPFlag("uniservice.txpool.nolocals", flags.Lookup("txpool_nolocals"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.TxPool.Journal,
		"txpool_journal",
		uniCfgInstance.UniServiceCfg.TxPool.Journal,
		"Disk journal for local transaction to survive node restarts",
	)
	viper.BindPFlag("uniservice.txpool.journal", flags.Lookup("txpool_journal"))

	flags.DurationVar(
		&uniCfgInstance.UniServiceCfg.TxPool.Rejournal,
		"txpool_rejournal",
		uniCfgInstance.UniServiceCfg.TxPool.Rejournal,
		"Time interval to regenerate the local transaction journal",
	)
	viper.BindPFlag("uniservice.txpool.rejournal", flags.Lookup("txpool_rejournal"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.PriceBump,
		"txpool_pricebump",
		uniCfgInstance.UniServiceCfg.TxPool.PriceBump,
		"Price bump percentage to replace an already existing transaction",
	)
	viper.BindPFlag("uniservice.txpool.pricebump", flags.Lookup("txpool_pricebump"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.PriceLimit,
		"txpool_pricelimit",
		uniCfgInstance.UniServiceCfg.TxPool.PriceLimit,
		"Minimum gas price limit to enforce for acceptance into the pool",
	)
	viper.BindPFlag("uniservice.txpool.pricelimit", flags.Lookup("txpool_pricelimit"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.AccountSlots,
		"txpool_accountslots",
		uniCfgInstance.UniServiceCfg.TxPool.AccountSlots,
		"Number of executable transaction slots guaranteed per account",
	)
	viper.BindPFlag("uniservice.txpool.accountslots", flags.Lookup("txpool_accountslots"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.AccountQueue,
		"txpool_accountqueue",
		uniCfgInstance.UniServiceCfg.TxPool.AccountQueue,
		"Maximum number of non-executable transaction slots permitted per account",
	)
	viper.BindPFlag("uniservice.txpool.accountqueue", flags.Lookup("txpool_accountqueue"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.GlobalSlots,
		"txpool_globalslots",
		uniCfgInstance.UniServiceCfg.TxPool.GlobalSlots,
		"Maximum number of executable transaction slots for all accounts",
	)
	viper.BindPFlag("uniservice.txpool.globalslots", flags.Lookup("txpool_globalslots"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.GlobalQueue,
		"txpool_globalqueue",
		uniCfgInstance.UniServiceCfg.TxPool.GlobalQueue,
		"Minimum number of non-executable transaction slots for all accounts",
	)
	viper.BindPFlag("uniservice.txpool.globalqueue", flags.Lookup("txpool_globalqueue"))

	flags.DurationVar(
		&uniCfgInstance.UniServiceCfg.TxPool.Lifetime,
		"txpool_lifetime",
		uniCfgInstance.UniServiceCfg.TxPool.Lifetime,
		"Maximum amount of time non-executable transaction are queued",
	)
	viper.BindPFlag("uniservice.txpool.lifetime", flags.Lookup("txpool_lifetime"))

	flags.DurationVar(
		&uniCfgInstance.UniServiceCfg.TxPool.ResendTime,
		"txpool_resendtime",
		uniCfgInstance.UniServiceCfg.TxPool.ResendTime,
		"Maximum amount of time  executable transaction are resended",
	)
	viper.BindPFlag("uniservice.txpool.resendtime", flags.Lookup("txpool_resendtime"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.MinBroadcast,
		"txpool_minbroadcast",
		uniCfgInstance.UniServiceCfg.TxPool.MinBroadcast,
		"Minimum number of nodes for the transaction broadcast",
	)
	viper.BindPFlag("uniservice.txpool.minbroadcast", flags.Lookup("txpool_minbroadcast"))

	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.TxPool.RatioBroadcast,
		"txpool_ratiobroadcast",
		uniCfgInstance.UniServiceCfg.TxPool.RatioBroadcast,
		"Ratio of nodes for the transaction broadcast",
	)
	viper.BindPFlag("uniservice.txpool.ratiobroadcast", flags.Lookup("txpool_ratiobroadcast"))

	// miner
	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.Miner.Start,
		"miner_start",
		uniCfgInstance.UniServiceCfg.Miner.Start,
		"Start miner generate block and process transaction",
	)
	viper.BindPFlag("uniservice.miner.start", flags.Lookup("miner_start"))

	// miner
	flags.Uint64Var(
		&uniCfgInstance.UniServiceCfg.Miner.Delay,
		"miner_delay",
		uniCfgInstance.UniServiceCfg.Miner.Delay,
		"delay duration for miner (ms)",
	)
	viper.BindPFlag("uniservice.miner.delay", flags.Lookup("miner_delay"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.Miner.Name,
		"miner_name",
		uniCfgInstance.UniServiceCfg.Miner.Name,
		"Name for block mining rewards",
	)
	viper.BindPFlag("uniservice.miner.name", flags.Lookup("miner_name"))

	flags.StringSliceVar(
		&uniCfgInstance.UniServiceCfg.Miner.PrivateKeys,
		"miner_private",
		uniCfgInstance.UniServiceCfg.Miner.PrivateKeys,
		"Hex of private key for block mining rewards",
	)
	viper.BindPFlag("uniservice.miner.private", flags.Lookup("miner_private"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.Miner.ExtraData,
		"miner_extra",
		uniCfgInstance.UniServiceCfg.Miner.ExtraData,
		"Block extra data set by the miner",
	)
	viper.BindPFlag("uniservice.miner.name", flags.Lookup("miner_extra"))

	// gas price oracle
	flags.IntVar(
		&uniCfgInstance.UniServiceCfg.GasPrice.Blocks,
		"gpo_blocks",
		uniCfgInstance.UniServiceCfg.GasPrice.Blocks,
		"Number of recent blocks to check for gas prices",
	)
	viper.BindPFlag("uniservice.gpo.blocks", flags.Lookup("gpo_blocks"))

	// metrics
	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.MetricsFlag,
		"metrics_start",
		uniCfgInstance.UniServiceCfg.MetricsConf.MetricsFlag,
		"flag that open statistical metrics",
	)
	viper.BindPFlag("uniservice.metrics.start", flags.Lookup("metrics_start"))

	flags.BoolVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.InfluxDBFlag,
		"metrics_influxdb",
		uniCfgInstance.UniServiceCfg.MetricsConf.InfluxDBFlag,
		"flag that open influxdb thad store statistical metrics",
	)
	viper.BindPFlag("uniservice.metrics.influxdb", flags.Lookup("metrics_influxdb"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.URL,
		"metrics_influxdb_URL",
		uniCfgInstance.UniServiceCfg.MetricsConf.URL,
		"URL that connect influxdb",
	)
	viper.BindPFlag("uniservice.metrics.influxdbURL", flags.Lookup("metrics_influxdb_URL"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.DataBase,
		"metrics_influxdb_name",
		uniCfgInstance.UniServiceCfg.MetricsConf.DataBase,
		"Influxdb database name",
	)
	viper.BindPFlag("uniservice.metrics.influxdbname", flags.Lookup("metrics_influxdb_name"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.UserName,
		"metrics_influxdb_user",
		uniCfgInstance.UniServiceCfg.MetricsConf.UserName,
		"Indluxdb user name",
	)
	viper.BindPFlag("uniservice.metrics.influxdbuser", flags.Lookup("metrics_influxdb_user"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.PassWd,
		"metrics_influxdb_passwd",
		uniCfgInstance.UniServiceCfg.MetricsConf.PassWd,
		"Influxdb user passwd",
	)
	viper.BindPFlag("uniservice.metrics.influxdbpasswd", flags.Lookup("metrics_influxdb_passwd"))

	flags.StringVar(
		&uniCfgInstance.UniServiceCfg.MetricsConf.NameSpace,
		"metrics_influxdb_namespace",
		uniCfgInstance.UniServiceCfg.MetricsConf.NameSpace,
		"Influxdb namespace",
	)
	viper.BindPFlag("uniservice.metrics.influxdbnamepace", flags.Lookup("metrics_influxdb_namespace"))

	// p2p
	flags.UintVar(
		&uniCfgInstance.NodeCfg.P2PConfig.NetworkID,
		"p2p_id",
		uniCfgInstance.NodeCfg.P2PConfig.NetworkID,
		"The ID of the p2p network. Nodes have different ID cannot communicate, even if they have same chainID and block data.",
	)
	viper.BindPFlag("uniservice.p2p.networkid", flags.Lookup("p2p_id"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PConfig.Name,
		"p2p_name",
		uniCfgInstance.NodeCfg.P2PConfig.Name,
		"The name sets the p2p node name of this server",
	)
	viper.BindPFlag("uniservice.p2p.name", flags.Lookup("p2p_name"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.P2PConfig.MaxPeers,
		"p2p_maxpeers",
		uniCfgInstance.NodeCfg.P2PConfig.MaxPeers,
		"Maximum number of network peers ",
	)
	viper.BindPFlag("uniservice.p2p.maxpeers", flags.Lookup("p2p_maxpeers"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.P2PConfig.MaxPendingPeers,
		"p2p_maxpendpeers",
		uniCfgInstance.NodeCfg.P2PConfig.MaxPendingPeers,
		"Maximum number of pending connection attempts ",
	)
	viper.BindPFlag("uniservice.p2p.maxpendpeers", flags.Lookup("p2p_maxpendpeers"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.P2PConfig.DialRatio,
		"p2p_dialratio",
		uniCfgInstance.NodeCfg.P2PConfig.DialRatio,
		"DialRatio controls the ratio of inbound to dialed connections",
	)
	viper.BindPFlag("uniservice.p2p.dialratio", flags.Lookup("p2p_dialratio"))

	flags.IntVar(
		&uniCfgInstance.NodeCfg.P2PConfig.PeerPeriod,
		"p2p_peerperiod",
		uniCfgInstance.NodeCfg.P2PConfig.PeerPeriod,
		"Disconnect the worst peer every 'p2p_peerperiod' ms(if peer count equal p2p_maxpeers), 0 means disable.",
	)
	viper.BindPFlag("uniservice.p2p.peerperiod", flags.Lookup("p2p_peerperiod"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PConfig.ListenAddr,
		"p2p_listenaddr",
		uniCfgInstance.NodeCfg.P2PConfig.ListenAddr,
		"Network listening address",
	)
	viper.BindPFlag("uniservice.p2p.listenaddr", flags.Lookup("p2p_listenaddr"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PNodeDatabase,
		"p2p_nodedb",
		uniCfgInstance.NodeCfg.P2PNodeDatabase,
		"The path to the database containing the previously seen live nodes in the network",
	)
	viper.BindPFlag("uniservice.p2p.nodedb", flags.Lookup("p2p_nodedb"))

	flags.BoolVar(
		&uniCfgInstance.NodeCfg.P2PConfig.NoDiscovery,
		"p2p_nodiscovery",
		uniCfgInstance.NodeCfg.P2PConfig.NoDiscovery,
		"Disables the peer discovery mechanism (manual peer addition)",
	)
	viper.BindPFlag("uniservice.p2p.nodiscovery", flags.Lookup("p2p_nodiscovery"))

	flags.BoolVar(
		&uniCfgInstance.NodeCfg.P2PConfig.NoDial,
		"p2p_nodial",
		uniCfgInstance.NodeCfg.P2PConfig.NoDial,
		"The server will not dial any peers.",
	)
	viper.BindPFlag("uniservice.p2p.nodial", flags.Lookup("p2p_nodial"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PBootNodes,
		"p2p_bootnodes",
		uniCfgInstance.NodeCfg.P2PBootNodes,
		"Node list file. BootstrapNodes are used to establish connectivity with the rest of the network",
	)
	viper.BindPFlag("uniservice.p2p.bootnodes", flags.Lookup("p2p_bootnodes"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PStaticNodes,
		"p2p_staticnodes",
		uniCfgInstance.NodeCfg.P2PStaticNodes,
		"Node list file. Static nodes are used as pre-configured connections which are always maintained and re-connected on disconnects",
	)
	viper.BindPFlag("uniservice.p2p.staticnodes", flags.Lookup("p2p_staticnodes"))

	flags.StringVar(
		&uniCfgInstance.NodeCfg.P2PTrustNodes,
		"p2p_trustnodes",
		uniCfgInstance.NodeCfg.P2PStaticNodes,
		"Node list file. Trusted nodes are usesd as pre-configured connections which are always allowed to connect, even above the peer limit",
	)
	viper.BindPFlag("uniservice.p2p.trustnodes", flags.Lookup("p2p_trustnodes"))

}
