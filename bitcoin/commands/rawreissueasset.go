// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
)

// RawReissueAsset
func RawReissueAsset(ctx context.Context,
	rpcClient RpcClient,
	hex Transaction,
	assetAmount float64,
	assetAddress, entropy, assetBlinding string,
	inputIndex int,
) (Transaction, error) {
	return rawReissueAssetWithOptions(ctx, rpcClient, hex, RawReissueAssetOptions{
		AssetAmount:  assetAmount,
		AssetAddress: assetAddress,
		Entropy:      entropy,
		AssetBlinder: assetBlinding,
		InputIndex:   inputIndex,
	})
}

func rawReissueAssetWithOptions(ctx context.Context,
	rpcClient RpcClient,
	hex Transaction,
	options RawReissueAssetOptions,
) (Transaction, error) {
	var result ReissuedTransaction
	var data []interface{}
	data = append(data, options)
	err := callCommand(rpcClient, CmdRawReissueAsset, &result, hex, &data)
	if err != nil {
		return result.Hex, err
	}
	return result.Hex, nil
}
