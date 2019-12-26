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
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/unichainplatform/unichain/accountmanager"
	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/crypto"
	testcommon "github.com/unichainplatform/unichain/test/common"
	"github.com/unichainplatform/unichain/types"
	"github.com/unichainplatform/unichain/utils/rlp"
)

var (
	privateKey, _ = crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	from          = common.Name("unichain.founder")
	to            = common.Name("unichain.account")
	aca           = common.Name("fcoinaccounta")
	acb           = common.Name("fcoinaccountb")
	acc           = common.Name("fcoinaccountc")
	acd           = common.Name("fcoinaccountd")

	a_author_0_priv *ecdsa.PrivateKey
	a_author_2_priv *ecdsa.PrivateKey
	a_author_3_priv *ecdsa.PrivateKey
	b_author_0_priv *ecdsa.PrivateKey
	b_author_2_priv *ecdsa.PrivateKey
	c_author_0_priv *ecdsa.PrivateKey
	c_author_1_priv *ecdsa.PrivateKey
	c_author_2_priv *ecdsa.PrivateKey
	d_author_0_priv *ecdsa.PrivateKey

	newPrivateKey_a *ecdsa.PrivateKey
	newPrivateKey_b *ecdsa.PrivateKey
	newPrivateKey_c *ecdsa.PrivateKey
	newPrivateKey_d *ecdsa.PrivateKey

	pubKey_a common.PubKey
	pubKey_b common.PubKey
	pubKey_c common.PubKey
	pubKey_d common.PubKey

	aNonce = uint64(0)
	bNonce = uint64(0)
	cNonce = uint64(0)
	dNonce = uint64(0)

	assetID  = uint64(0)
	nonce    = uint64(0)
	gasLimit = uint64(2000000)
)

func generateAccount() {
	nonce, _ = testcommon.GetNonce(from)

	newPrivateKey_a, _ = crypto.GenerateKey()
	pubKey_a = common.BytesToPubKey(crypto.FromECDSAPub(&newPrivateKey_a.PublicKey))
	a_author_0_priv = newPrivateKey_a
	fmt.Println("priv_a ", hex.EncodeToString(crypto.FromECDSA(newPrivateKey_a)), " pubKey_a ", pubKey_a.String())

	newPrivateKey_b, _ = crypto.GenerateKey()
	pubKey_b = common.BytesToPubKey(crypto.FromECDSAPub(&newPrivateKey_b.PublicKey))
	b_author_0_priv = newPrivateKey_b
	fmt.Println("priv_b ", hex.EncodeToString(crypto.FromECDSA(newPrivateKey_b)), " pubKey_b ", pubKey_b.String())

	newPrivateKey_c, _ = crypto.GenerateKey()
	pubKey_c = common.BytesToPubKey(crypto.FromECDSAPub(&newPrivateKey_c.PublicKey))
	c_author_0_priv = newPrivateKey_c
	fmt.Println("priv_c ", hex.EncodeToString(crypto.FromECDSA(newPrivateKey_c)), " pubKey_c ", pubKey_c.String())

	newPrivateKey_d, _ = crypto.GenerateKey()
	pubKey_d = common.BytesToPubKey(crypto.FromECDSAPub(&newPrivateKey_d.PublicKey))
	d_author_0_priv = newPrivateKey_d
	fmt.Println("priv_d ", hex.EncodeToString(crypto.FromECDSA(newPrivateKey_d)), " pubKey_d ", pubKey_d.String())

	balance, _ := testcommon.GetAccountBalanceByID(from, assetID)
	balance.Div(balance, big.NewInt(10))

	aca = common.Name(fmt.Sprintf("fcoinaccounta%d", nonce))
	acb = common.Name(fmt.Sprintf("fcoinaccountb%d", nonce))
	acc = common.Name(fmt.Sprintf("fcoinaccountc%d", nonce))
	acd = common.Name(fmt.Sprintf("fcoinaccountd%d", nonce))

	key := types.MakeKeyPair(privateKey, []uint64{0})
	acct := &accountmanager.CreateAccountAction{
		AccountName: aca,
		Founder:     aca,
		PublicKey:   pubKey_a,
	}
	b, _ := rlp.EncodeToBytes(acct)
	sendTransferTx(types.CreateAccount, from, to, nonce, assetID, balance, b, []*types.KeyPair{key}, nil, nil)

	acct = &accountmanager.CreateAccountAction{
		AccountName: acb,
		Founder:     acb,
		PublicKey:   pubKey_b,
	}
	b, _ = rlp.EncodeToBytes(acct)
	sendTransferTx(types.CreateAccount, from, to, nonce+1, assetID, balance, b, []*types.KeyPair{key}, nil, nil)

	acct = &accountmanager.CreateAccountAction{
		AccountName: acc,
		Founder:     acc,
		PublicKey:   pubKey_c,
	}
	b, _ = rlp.EncodeToBytes(acct)
	sendTransferTx(types.CreateAccount, from, to, nonce+2, assetID, balance, b, []*types.KeyPair{key}, nil, nil)

	acct = &accountmanager.CreateAccountAction{
		AccountName: acd,
		Founder:     acd,
		PublicKey:   pubKey_d,
	}
	b, _ = rlp.EncodeToBytes(acct)
	sendTransferTx(types.CreateAccount, from, to, nonce+3, assetID, balance, b, []*types.KeyPair{key}, nil, nil)

	for {
		time.Sleep(10 * time.Second)
		aexist, _ := testcommon.CheckAccountIsExist(aca)
		bexist, _ := testcommon.CheckAccountIsExist(acb)
		cexist, _ := testcommon.CheckAccountIsExist(acc)
		dexist, _ := testcommon.CheckAccountIsExist(acd)

		acaAccount, _ := testcommon.GetAccountByName(aca)
		acbAccount, _ := testcommon.GetAccountByName(acb)
		accAccount, _ := testcommon.GetAccountByName(acc)
		acdAccount, _ := testcommon.GetAccountByName(acd)

		fmt.Println("acaAccount version hash", acaAccount.AuthorVersion.Hex())
		fmt.Println("acbAccount version hash", acbAccount.AuthorVersion.Hex())
		fmt.Println("accAccount version hash", accAccount.AuthorVersion.Hex())
		fmt.Println("accAccount version hash", acdAccount.AuthorVersion.Hex())

		if aexist && bexist && cexist && dexist {
			break
		}
	}

	fmt.Println("aca ", aca, " acb ", acb, " acc ", acc, " acd ", acd)
}

func init() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	generateAccount()
}

func addAuthorsForAca() {
	a_author_0 := common.NewAuthor(pubKey_a, 500)
	a_authorAct_0 := &accountmanager.AuthorAction{1, a_author_0}

	a_author_1 := common.NewAuthor(acb, 400)
	a_authorAct_1 := &accountmanager.AuthorAction{0, a_author_1}

	a_author_2_priv, _ = crypto.GenerateKey()
	a_author_2_addr := crypto.PubkeyToAddress(a_author_2_priv.PublicKey)
	a_author_2 := common.NewAuthor(a_author_2_addr, 400)
	a_authorAct_2 := &accountmanager.AuthorAction{0, a_author_2}

	a_author_3_priv, _ = crypto.GenerateKey()
	a_author_3_pub := common.BytesToPubKey(crypto.FromECDSAPub(&a_author_3_priv.PublicKey))
	a_author_3 := common.NewAuthor(a_author_3_pub, 400)
	a_authorAct_3 := &accountmanager.AuthorAction{0, a_author_3}

	authorAction := make([]*accountmanager.AuthorAction, 0)
	authorAction = append(authorAction, a_authorAct_0, a_authorAct_1, a_authorAct_2, a_authorAct_3)

	action := &accountmanager.AccountAuthorAction{1000, 0, authorAction}
	input, err := rlp.EncodeToBytes(action)
	if err != nil {
		jww.INFO.Println("addAuthors for accounta error ... ", err)
		return
	}
	key := types.MakeKeyPair(newPrivateKey_a, []uint64{0})

	sendTransferTx(types.UpdateAccountAuthor, aca, to, aNonce, assetID, big.NewInt(0), input, []*types.KeyPair{key}, nil, nil)
}

func addAuthorsForAcb() {
	b_author_0 := common.NewAuthor(pubKey_b, 50)
	b_authorAct_0 := &accountmanager.AuthorAction{1, b_author_0}

	b_author_1 := common.NewAuthor(acc, 40)
	b_authorAct_1 := &accountmanager.AuthorAction{0, b_author_1}

	b_author_2_priv, _ = crypto.GenerateKey()
	b_author_2_addr := crypto.PubkeyToAddress(b_author_2_priv.PublicKey)
	b_author_2 := common.NewAuthor(b_author_2_addr, 40)
	b_authorAct_2 := &accountmanager.AuthorAction{0, b_author_2}

	action := &accountmanager.AccountAuthorAction{100, 0, []*accountmanager.AuthorAction{b_authorAct_0, b_authorAct_1, b_authorAct_2}}
	input, err := rlp.EncodeToBytes(action)
	if err != nil {
		jww.INFO.Println("addAuthors for accountb error ... ", err)
		return
	}
	key := types.MakeKeyPair(newPrivateKey_b, []uint64{0})
	sendTransferTx(types.UpdateAccountAuthor, acb, to, bNonce, assetID, big.NewInt(0), input, []*types.KeyPair{key}, nil, nil)
}

func addAuthorsForAcc() {
	c_author_0 := common.NewAuthor(pubKey_c, 5)
	c_authorAct_0 := &accountmanager.AuthorAction{1, c_author_0}

	c_author_1_priv, _ = crypto.GenerateKey()
	c_author_1_addr := crypto.PubkeyToAddress(c_author_1_priv.PublicKey)
	c_author_1 := common.NewAuthor(c_author_1_addr, 4)
	c_authorAct_1 := &accountmanager.AuthorAction{0, c_author_1}

	c_author_2_priv, _ = crypto.GenerateKey()
	c_author_2_pub := common.BytesToPubKey(crypto.FromECDSAPub(&c_author_2_priv.PublicKey))
	c_author_2 := common.NewAuthor(c_author_2_pub, 4)
	c_authorAct_2 := &accountmanager.AuthorAction{0, c_author_2}

	action := &accountmanager.AccountAuthorAction{10, 0, []*accountmanager.AuthorAction{c_authorAct_0, c_authorAct_1, c_authorAct_2}}
	input, err := rlp.EncodeToBytes(action)
	if err != nil {
		jww.INFO.Println("addAuthors for accountc error ... ", err)
		return
	}
	key := types.MakeKeyPair(newPrivateKey_c, []uint64{0})
	sendTransferTx(types.UpdateAccountAuthor, acc, to, cNonce, assetID, big.NewInt(0), input, []*types.KeyPair{key}, nil, nil)
}

func transferFromA2B() {
	key_1_0 := types.MakeKeyPair(b_author_0_priv, []uint64{1, 0})
	key_1_1_0 := types.MakeKeyPair(c_author_0_priv, []uint64{1, 1, 0})
	key_1_1_1 := types.MakeKeyPair(c_author_1_priv, []uint64{1, 1, 1})
	key_1_1_2 := types.MakeKeyPair(c_author_2_priv, []uint64{1, 1, 2})
	key_2 := types.MakeKeyPair(a_author_2_priv, []uint64{2})
	key_3 := types.MakeKeyPair(a_author_3_priv, []uint64{3})
	key_1_2 := types.MakeKeyPair(b_author_2_priv, []uint64{1, 2})

	aNonce++
	sendTransferTx(types.Transfer, aca, to, aNonce, assetID, big.NewInt(1), nil, []*types.KeyPair{key_1_0, key_1_1_0, key_1_1_1, key_1_1_2, key_2, key_3, key_1_2}, nil, nil)
}

func modifyAUpdateAUthorThreshold() {
	key_1_0 := types.MakeKeyPair(b_author_0_priv, []uint64{1, 0})
	key_1_1_0 := types.MakeKeyPair(c_author_0_priv, []uint64{1, 1, 0})
	key_1_1_1 := types.MakeKeyPair(c_author_1_priv, []uint64{1, 1, 1})
	key_1_1_2 := types.MakeKeyPair(c_author_2_priv, []uint64{1, 1, 2})
	key_2 := types.MakeKeyPair(a_author_2_priv, []uint64{2})
	key_3 := types.MakeKeyPair(a_author_3_priv, []uint64{3})
	key_1_2 := types.MakeKeyPair(b_author_2_priv, []uint64{1, 2})
	key_0 := types.MakeKeyPair(a_author_0_priv, []uint64{0})

	action := &accountmanager.AccountAuthorAction{0, 2, []*accountmanager.AuthorAction{}}
	input, err := rlp.EncodeToBytes(action)
	if err != nil {
		jww.INFO.Println("addAuthors for accountc error ... ", err)
		return
	}

	aNonce++
	sendTransferTx(types.UpdateAccountAuthor, aca, to, aNonce, assetID, big.NewInt(0), input, []*types.KeyPair{key_1_0, key_1_1_0, key_1_1_1, key_1_1_2, key_2, key_3, key_1_2, key_0}, nil, nil)
}

func transferFromA2BWithBAsPayer() {
	key_1_0 := types.MakeKeyPair(b_author_0_priv, []uint64{1, 0})
	key_1_1_0 := types.MakeKeyPair(c_author_0_priv, []uint64{1, 1, 0})
	key_1_1_1 := types.MakeKeyPair(c_author_1_priv, []uint64{1, 1, 1})
	key_1_1_2 := types.MakeKeyPair(c_author_2_priv, []uint64{1, 1, 2})
	key_2 := types.MakeKeyPair(a_author_2_priv, []uint64{2})
	key_3 := types.MakeKeyPair(a_author_3_priv, []uint64{3})
	key_1_2 := types.MakeKeyPair(b_author_2_priv, []uint64{1, 2})

	gasPrice, _ := testcommon.GasPrice()
	fp := &types.FeePayer{
		GasPrice: gasPrice,
		Payer:    acd,
		Sign:     &types.Signature{0, make([]*types.SignData, 0)},
	}
	payerKey := types.MakeKeyPair(newPrivateKey_d, []uint64{0})

	aNonce++
	sendTransferTx(types.Transfer, aca, to, aNonce, assetID, big.NewInt(1), nil, []*types.KeyPair{key_1_0, key_1_1_0, key_1_1_1, key_1_1_2, key_2, key_3, key_1_2}, fp, []*types.KeyPair{payerKey})
}

func main() {
	jww.INFO.Println("test send sundry transaction...")

	// d_author_1_priv, _ := crypto.GenerateKey()
	// d_author_1_pub := common.BytesToPubKey(crypto.FromECDSAPub(&d_author_1_priv.PublicKey))
	// d_author_1_pub_addr := common.BytesToAddress(crypto.Keccak256(d_author_1_pub.Bytes()[1:])[12:])
	// d_author_1_addr := crypto.PubkeyToAddress(d_author_1_priv.PublicKey)
	// fmt.Println(d_author_1_pub_addr.String(), d_author_1_addr.String())

	addAuthorsForAca()
	addAuthorsForAcb()
	addAuthorsForAcc()
	time.Sleep(10 * time.Second)

	acaAccount, _ := testcommon.GetAccountByName(aca)
	acbAccount, _ := testcommon.GetAccountByName(acb)
	accAccount, _ := testcommon.GetAccountByName(acc)
	fmt.Println("update acaAccount version hash", acaAccount.AuthorVersion.Hex())
	fmt.Println("update acbAccount version hash", acbAccount.AuthorVersion.Hex())
	fmt.Println("update accAccount version hash", accAccount.AuthorVersion.Hex())

	transferFromA2B()
	modifyAUpdateAUthorThreshold()

	transferFromA2BWithBAsPayer()
}

func sendTransferTx(txType types.ActionType, from, to common.Name, nonce, assetID uint64, value *big.Int, input []byte, keys []*types.KeyPair, fp *types.FeePayer, payerKeys []*types.KeyPair) {
	action := types.NewAction(txType, from, to, nonce, assetID, gasLimit, value, input, nil)
	gasprice, _ := testcommon.GasPrice()
	if fp != nil {
		gasprice = big.NewInt(0)
	}
	tx := types.NewTransaction(0, gasprice, action)

	signer := types.MakeSigner(big.NewInt(1))
	err := types.SignActionWithMultiKey(action, tx, signer, 0, keys)
	if err != nil {
		jww.ERROR.Fatalln(err)
	}

	if fp != nil {
		err = types.SignPayerActionWithMultiKey(action, tx, signer, fp, 0, payerKeys)
		if err != nil {
			jww.ERROR.Fatalln(err)
		}
	}

	rawtx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		jww.ERROR.Fatalln(err)
	}

	hash, err := testcommon.SendRawTx(rawtx)
	if err != nil {
		jww.INFO.Println("result err: ", err)

	}
	jww.INFO.Println("result hash: ", hash.Hex())
}
