// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package common

import (
	"context"

	"github.com/condensat/bank-wallet/ssm/commands"
)

// ChainClient interface specification for bitcoin and elements
type ChainClient interface {
	GetNewAddress(ctx context.Context, account string) (string, error)
	GetAddressInfo(ctx context.Context, address string) (AddressInfo, error)
	GetBlockCount(ctx context.Context) (int64, error)
	ListUnspent(ctx context.Context, minConf, maxConf int, addresses ...string) ([]TransactionInfo, error)
	LockUnspent(ctx context.Context, unlock bool, transactions ...TransactionInfo) error
	ListLockUnspent(ctx context.Context) ([]TransactionInfo, error)
	GetTransaction(ctx context.Context, txID string) (TransactionInfo, error)

	SpendFunds(ctx context.Context, inputs []UTXOInfo, outputs []SpendInfo) (SpendTx, error)
}

// SsmClient interface specification for crypto-ssm
type SsmClient interface {
	NewAddress(ctx context.Context, ssmPath commands.SsmPath) (string, error)
	SignTx(ctx context.Context, chain, inputransaction string, inputs ...commands.SignTxInputs) (string, error)
}
