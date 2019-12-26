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

package rpcapi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/unichainplatform/unichain/common"
	"github.com/unichainplatform/unichain/types"
	"github.com/unichainplatform/unichain/utils/rlp"
)

// PublicUniChainAPI offers and API for the transaction pool. It only operates on data that is non confidential.
type PublicUniChainAPI struct {
	b Backend
}

// NewPublicUniChainAPI creates a new tx pool service that gives information about the transaction pool.
func NewPublicUniChainAPI(b Backend) *PublicUniChainAPI {
	return &PublicUniChainAPI{b}
}

// GasPrice returns a suggestion for a gas price.
func (s *PublicUniChainAPI) GasPrice(ctx context.Context) (*big.Int, error) {
	return s.b.SuggestPrice(ctx)
}

// SendRawTransaction will add the signed transaction to the transaction pool.
// The sender is responsible for signing the transaction and using the correct nonce.
func (s *PublicUniChainAPI) SendRawTransaction(ctx context.Context, encodedTx hexutil.Bytes) (common.Hash, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(encodedTx, tx); err != nil {
		return common.Hash{}, err
	}
	return submitTransaction(ctx, s.b, tx)
}
