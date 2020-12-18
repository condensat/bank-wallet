// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
)

// RawIssueAssetWithAsset This is the minimal number of arguments you need to pass to issue an asset
func RawIssueAssetWithAsset(ctx context.Context, rpcClient RpcClient, hex Transaction, assetAddress string, assetAmount float64) (IssuedTransaction, error) {
	return rawIssueAssetWithOptions(ctx, rpcClient, hex, RawIssueAssetOptions{
		AssetAmount:  assetAmount,
		AssetAddress: assetAddress,
		Blind:        true, //we suppose that we always want to blind issuance, maybe we should change it though
	})
}

func RawIssueAssetWithToken(ctx context.Context, rpcClient RpcClient, hex Transaction, assetAddress string, assetAmount float64, tokenAddress string, tokenAmount float64) (IssuedTransaction, error) {
	return rawIssueAssetWithOptions(ctx, rpcClient, hex, RawIssueAssetOptions{
		AssetAmount:  assetAmount,
		AssetAddress: assetAddress,
		TokenAmount:  tokenAmount,
		TokenAddress: tokenAddress,
		Blind:        true,
	})
}

func RawIssueAssetWithContract(ctx context.Context, rpcClient RpcClient, hex Transaction, assetAddress string, assetAmount float64, contractHash string) (IssuedTransaction, error) {
	return rawIssueAssetWithOptions(ctx, rpcClient, hex, RawIssueAssetOptions{
		AssetAmount:  assetAmount,
		AssetAddress: assetAddress,
		ContractHash: contractHash, //this is 64B long
		Blind:        true,
	})
}

func RawIssueAssetWithTokenWithContract(ctx context.Context, rpcClient RpcClient, hex Transaction, assetAddress string, assetAmount float64, tokenAddress string, tokenAmount float64, contractHash string) (IssuedTransaction, error) {
	return rawIssueAssetWithOptions(ctx, rpcClient, hex, RawIssueAssetOptions{
		AssetAmount:  assetAmount,
		AssetAddress: assetAddress,
		TokenAmount:  tokenAmount,
		TokenAddress: tokenAddress,
		ContractHash: contractHash, //this is 64B long
		Blind:        true,
	})
}

func rawIssueAssetWithOptions(ctx context.Context, rpcClient RpcClient, hex Transaction, options RawIssueAssetOptions) (IssuedTransaction, error) {
	var result []IssuedTransaction
	var data []interface{}
	data = append(data, options)
	err := callCommand(rpcClient, CmdRawIssueAsset, &result, hex, &data)
	if err != nil {
		return IssuedTransaction{}, err
	}
	return result[0], nil
}
