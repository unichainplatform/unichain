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

package accountmanager

import (
	"math/big"

	"github.com/unichainplatform/unichain/asset"
	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/types"
)

//IAccount export account interface
type IAccount interface {
	GetName() common.Name
	//nonce
	GetNonce() uint64
	SetNonce(nonce uint64)
	//code
	GetCode() ([]byte, error)
	SetCode(code []byte) (bool, error)
	GetCodeHash() (common.Hash, error)
	GetCodeSize() (uint64, error)
	//
	GetPubKey() common.PubKey
	SetPubKey(pubkey common.PubKey)
	//asset
	GetBalancesList() ([]*AssetBalance, error)
	GetAllAccountBalance() (map[uint64]*big.Int, error)
	AddBalanceByID(assetID uint64, value *big.Int) error
	SubBalanceByID(assetID uint64, value *big.Int) error
	EnoughAccountBalance(assetID uint64, value *big.Int) error
	//
	IsSuicided() bool
	SetSuicide()
	//
	IsDestroyed()
	SetDestroy()
}

//IAccountManager export account manager interface
type IAccountManager interface {
	//account
	AccountIsExist(accountName common.Name) (bool, error)
	AccountIsEmpty(accountName common.Name) (bool, error)
	CreateAccount(accountName common.Name, pubkey common.PubKey) error
	DeleteAccountByName(accountName common.Name) error
	GetAccountByName(accountName common.Name) (*Account, error)
	SetAccount(acct *Account) error
	//sign
	RecoverTx(signer types.Signer, tx *types.Transaction) error
	IsValidSign(accountName common.Name, aType types.ActionType, pub common.PubKey) error
	//asset
	IssueAsset(asset *asset.AssetObject) error
	IncreaseAsset(accountName common.Name, assetID uint64, amount *big.Int) error
	//
	CanTransfer(accountName common.Name, assetId uint64, value *big.Int) (bool, error)
	TransferAsset(fromAccount common.Name, toAccount common.Name, assetID uint64, value *big.Int) error
	IncAsset2Acct(fromName common.Name, toName common.Name, assetId uint64, amount *big.Int) error
	AddBalanceByName(accountName common.Name, assetID uint64, amount *big.Int) error
	Process(action *types.Action) error
}

//IStateDB  export state database interface
type IStateDB interface {
	Put(account string, key string, value []byte)
	Get(account string, key string) ([]byte, error)
	RevertToSnapshot(revisionID int)
}
