// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"errors"

	"github.com/condensat/bank-wallet/rpc"
)

var (
	ErrInvalidRPCClient = errors.New("Invalid RPC Client")
)

func GetBlockCount(ctx context.Context, rpcClient RpcClient) (int64, error) {
	if rpcClient == nil {
		return 0, ErrInvalidRPCClient
	}

	var blockount int64
	err := callCommand(rpcClient, CmdGetBlockCount, &blockount)
	if err != nil {
		return 0, rpc.ErrRpcError
	}

	return blockount, nil
}