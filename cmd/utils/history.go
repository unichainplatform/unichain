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

package utils

import "github.com/monax/relic"

// Use below as template for change notes, delete empty sections but keep order
/*
### Security

### Changed

### Fixed

### Added

### Removed

### Deprecated

### Forked
*/

// History the releases described by version string and changes, newest release first.
// The current release is taken to be the first release in the slice, and its
// version determines the single authoritative version for the next release.
//
// To cut a new release add a release to the front of this slice then run the
// release tagging script: ./scripts/tag_release.sh
var History relic.ImmutableHistory = relic.NewHistory("unichain", "https://github.com/unichainplatform/unichain").
	MustDeclareReleases(
		"1.0.0 - 2019-12-09",
		`### Forked
- [DPOS] allow contract asset transfer (#525)(#528)
- [FEE] other people pay transaction fee (#531)(#533)(#536)
### Fixed
- [FEE]fee transfer internal record (#495)
- [BLOCKCHIAN] fixed export blockchain error (#498)
- [GAS] modify gas price (#501)
- [MINER] add setcoinbase check (#500) and fix miner bug (#499)(#511)(#512)(#513)(#514)(#516)
- [P2P] fixed bug that may close a nil channel (#503)and fixed ddos check error (#519)
- [DOWNLOAD]add node into blacklist if it had too much errors(#519)(#523)
### Added
- [CMD] add version cmd compile date info (#505)(#521)
- [CMD] cmd/uni: add method 'seednodes' into sub-cmd 'p2p' (#497)
- [CMD] add txpool cmd gettxsbyaccount (#502)
- [P2P] p2p,rpc: add rpc to query seed nodes from db(#496)
- [TEST] add each code module unit test or note (#492)(#493)(#504)(#508)(#507)(#509)
`,
		"0.0.30 - 2019-09-20",
		`### Fixed
- [DOWNLOADER] blockchain/downloader: if there were two different blocks in same height, it may cause download error (#481)(#490)
- [BLOCKCHAIN] blockchain: fixed bug that use nanosecond compared with second (#474)
### Changed
- [RAWDB] modify unichain db interface (#487)
- [BLOCKCHAIN] modify the blockchain code (#489)
### Added
- [GASPRICE] returns the recommended gas price (#486)
- [TEST] add each code module unit test or note (#471)(#472)(#473)(#475)(#476)(#478)(#482)(#483)(#485)
`,
		"0.0.29 - 2019-09-04",
		`### Fixed
- [P2P] blockchain/downloader: fixed bug that may don't download the new block (#469)
### Added
- [CMD] add chain command forkstatus (#468)
- [TEST] add each code module unit test or note.(#467)(#466)
`,
		"0.0.28 - 2019-09-04",
		`### Fixed
- [P2P] bootnode in genesis is uni finder.(#439)
- [DPOS]dpos api epoch add url. (#437)
- [INTERNALTX]modify internal tx log hash. (#441)
- [FILTER]fix rpc output format. (#440)
### Changed
- [RPCAPI] modify setStatePruning rpc private (#464),modify blockchain misspelling and change uni_getTransactions (#457)
### Added
- [RPCAPI] add rpc api uni_forkStatus.(#463)
- [TEST] add each code module unit test or note.(#442)(#446)(#447)(#448)(#450)(#451)(#454)(#455)(#456)(#458)(#459)(#461)(#462)(#465)
`,
		"0.0.27 - 2019-08-16",
		`### Fixed
- [VIM] vm fix stack valid (#416)
### Added
- [RPC] add rpc interface: uni_getTransBatch
`,
		"0.0.26 - 2019-08-14",
		`### Fixed
- [LOG] fix vendor log bug, patch for version 0.0.25 
`,
		"0.0.25 - 2019-08-07",
		`### Forked
- [DPOS] fork3: reduce CandidateAvailableMinQuantity (#416)
### Fixed
- [MINER] fixed some bugs (#421)(#422)
- [TXPOOL] fixed txpool add remotes not check the same tx was exist (#423)
- [TXPOOL] fixed txpool rpc send same tx err (#430)
- [ACCOUNT] fix internal action bug (#424)
### Changed
- [RPCAPI] modify api get account by name and id (#428)
- [GASPRICE] modify gas price oracle (#417)
- [ACCOUNT] modify account fliter balance zero (#414)
### Added
- [FILTERS] add filters rpc (#431)
- [MOD] support go mod (#429)
- [VM] add timeout func (#419) and add callwithpay (#425)
`,
		"0.0.24 - 2019-07-30",
		`### Fixed
- [BLOCKCHAIN] blockchain store irreversible number 
- [TXPOOL] fixed txpool test failed in travis CI
### Changed
- [DPOS] update some dpos apis
### Added
- [LOG] add some log print
- [RPC] add rpc dpos_snapShotStake and fixed GetActivedCandidate
`,
		"0.0.23 - 2019-07-15",
		`### Fixed
- [RPC] fixed getTxsByAccount rpc arg check and uint infinite loop
- [BLOCKCHAIN] modify blockchain start err 
### Changed
- [TXPOOL] move TxPool reorg and events to background goroutine
- [P2P] unifinder: add cmd flag that can input genesis block hash
### Added
- [P2P] txpool.handler: add config of txs broadcast
- [RPC] add some dpos rpc api for browser
`,
		"0.0.22 - 2019-06-24",
		`### Forked
- [ACCOUNTNAME] forkID=1: modify account verification rules,asset contains account prefix
### Changed
- [DPOS] modify dpos getepoch api
- [GENESIS] fix SetupGenesisBlock func return result
### Added
- [GENESIS] start node with fork id
`,
		"0.0.21 - 2019-06-15",
		`### Fixed
- [DOWNLOADER] fixed bug that may casue dead loop 
- [BLOCKCHAIN] fixed state store irreversible number bug
- [DPOS] fixed replace rate for candiate
### Removed
- [TXPOOL] removed some unused variable in txpool/handler.go
- [RPC] removed invalid code
### Added
- [TXPOOL] limited the amount of gorouting not greater 1024
- [GENESIS] add use default block gaslimit and update genesis.json
`,
		"0.0.20 - 2019-06-12",
		`### Fixed
- [DOWNLOADER] fixed bug of find ancestor and use random station 
- [BLOCKCHAIN] fixed blockchain irreversible number
### Add
- [DPOS] add thread test for rand vote candidate
- [BLOCKCHAIN] add refuse bad block hashes
- [BLOCKCHAIN] sync block with a specified block number
`,
		"0.0.19 - 2019-06-11",
		`### Fixed
- [ASSET] modify subasset decimals
`,
		"0.0.18 - 2019-06-06",
		`### Fixed
- [ACCOUNT] modify children check function
### Add
- [CONTRACT] contract add getassetid api 
- [MINER] fix should counter & add delay duration for miner
`,
		"0.0.17 - 2019-06-05",
		`### Changed
- [GENESIS] modify blockchain sys account name
### Fixed
- [BLOCKCHAIN] modify blockchain.HasState function
- [RPC] fix GetDelegatedByTime rpc interface
`,
		"0.0.16 - 2019-06-04",
		`### Changed
- [MAKEFILE] fixed bug of target build_workspace
- [ACCOUNT] account author lenght should not exceed 10
- [VM] modify gas distribution
### Add
- [DPOS] add min available quantity of candidate for vote 
- [CMD] add read yaml and toml test 
- [SDK] add sdk contract test
- [TYPES] support parentIndex when sign
- [TXPOOL] add txpool resend pending txs
### Fixed
- [P2P] broadcast txs to atleast 3 peers 
- [BLOCKCHAIN] downloader disconnected peers which has to much wrong blocks
- [DPOS] fix calc should counter of candidate
- [ALL] fixs some bugs
`,
		"0.0.15 - 2019-05-21",
		`### Changed
- [VM] change withdraw type to transfer
### Add
- [P2P] add flow control,some quit channel
- [P2p] periodic remove the worst peer if peer connections is full, but default is disabled.
- [RPC] add dpos rpc api for info by epcho
### Fixed
- [DPOS] fix bug when dpos started
- [ALL] fixs some bugs
`,
		"0.0.14 - 2019-05-20",
		`### Fixed
- [GENESIS] fix genesis bootnodes prase failed not start node
`,
		"0.0.13 - 2019-05-18",
		`### Add
- [GPO] add add gas price oracle unit test 
- [VM] move gas to GasTableInstance
### Fixed
- [PARAMS] change genesis gas limit to 30 million 
- [VM] opCreate doing nothing but push zero into stack and distributeGasByScale distribute right num
- [ACCOUNT] add check asset contract name, check account name length 
- [ALL] fixs some bugs
`,
		"0.0.12 - 2019-05-13",
		`### Add
- [CMD] add p2p miner txpool command.
### Deprecated
- [RPCAPI] modify account and blockchain return result
- [DOC] add jsonrpc, cmd, p2p docs in wiki
`,
		"0.0.11 - 2019-05-06",
		`### Deprecated
- [ASSET] modify asset and account action struct
- [ACCOUNT] modify account detail to description
- [DPOS] add dpos reward interface for solidity
`,
		"0.0.10 - 2019-05-06",
		`### Fixed
- [ASSET] get asset object return panic
- [VM] fix contract issue asset bug
`,
		"0.0.9 - 2019-05-06",
		`### Added
- [BLOCKCHAIN] add gensis block account
- [FEE] the distributed gas will add to unichain.fee's balance
- [COMMON] add json unmarshal for author
- [ASSET] check valid for modifing about contract asset
### Fixed
- [VM] execWithdrawFee return err when fm.WithdrawFeeFromSystem fail
- [BLOCKCHAIN] fix fork contracl init err
- [GENESIS] genesis block action repeat
- [DPOS] fix updateElectedCandidates bug when dpos is false
- [ALL] fixs some bugs
### Changed
- [COMMON] modify name for support more scenes and modify subaccount/subasset name
- [ASSET] modify issue asset return assetID
`,
		"0.0.8 - 2019-04-30",
		`### Added
- [DEBUG] add debug pprof,trace cmd flags and rpc
- [FEE] add fee manager and some rpc interface
- [TXPOOL] add bloom in transaction P2P message
- [TYPES] types/action.go add remark field
### Fixed
- [TXPOOL] fixed txpool queue and pending don't remove no permissions transactions
- [VM] fix bug that distribute more gas than given when internal call happens
- [BLOCKCHAIN] fixed restart node missmatch genesis block hash
- [ACCOUNTMANAGER] generate author version when account create
- [DPOS] solve infinite loop for getvoters
- [ALL] fixs some bugs
`,
		"0.0.7 - 2019-04-23",
		`### Removed
- [WALLET] removed wallet module，the local node not support store private key
### Added
- [VM] add opt for opSnapBalance and destroyasset for contract
- [BLOCKCHAIN] support import/export block
- [RPC] add get the contract internal transaction
### Fixed
- [VM] add opt for opSnapBalance
- [TYPES] fixs the base types
- [ALL] fixs some bugs
`,
		"0.0.6 - 2019-04-04",
		`### Added
- [CRYPTO] add btcd secp256k1 crypto
### Fixed
- [MAKEFILE] fixed cross platform
`,
		"0.0.5 - 2019-04-04",
		`### Added
- [README] add license badge
- [SCRIPTS] add is_checkout_dirty.sh release.sh tag_release.sh commit_hash.sh
### Fixed
- [MAKEFILE] add check fmt tag_release release command
`,
	)
