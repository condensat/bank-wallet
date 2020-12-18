// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"errors"

	"github.com/condensat/bank-core/appcontext"
	"github.com/condensat/bank-core/logger"
	"github.com/condensat/bank-core/messaging"
	"github.com/condensat/bank-wallet/common"
)

const (
	DefaultConfidentialIssuance = true
)

// AssetIssuance issues a new asset without reissuance token nor contract hash
func AssetIssuance(ctx context.Context, chain string, issuerID uint64, assetAddress string, assetAmount float64) (common.IssuanceResponse, error) {
	return assetIssuance(ctx, chain, DefaultConfidentialIssuance, issuerID, common.IssuanceRequest{
		Mode:               common.AssetIssuanceModeWithAsset,
		AssetPublicAddress: assetAddress,
		AssetIssuedAmount:  assetAmount,
	})
}

func AssetIssuanceWithToken(ctx context.Context, chain string, issuerID uint64, assetAddress string, assetAmount float64, tokenAddress string, tokenAmount float64) (common.IssuanceResponse, error) {
	return assetIssuance(ctx, chain, DefaultConfidentialIssuance, issuerID, common.IssuanceRequest{
		Mode:               common.AssetIssuanceModeWithToken,
		AssetPublicAddress: assetAddress,
		AssetIssuedAmount:  assetAmount,
		TokenPublicAddress: tokenAddress,
		TokenIssuedAmount:  tokenAmount,
	})
}

func AssetIssuanceWithContract(ctx context.Context, chain string, issuerID uint64, assetAddress string, assetAmount float64, contractHash string) (common.IssuanceResponse, error) {
	return assetIssuance(ctx, chain, DefaultConfidentialIssuance, issuerID, common.IssuanceRequest{
		Mode:               common.AssetIssuanceModeWithContract,
		AssetPublicAddress: assetAddress,
		AssetIssuedAmount:  assetAmount,
		ContractHash:       contractHash,
	})
}

func AssetIssuanceWithTokenWithContract(ctx context.Context, chain string, issuerID uint64, assetAddress string, assetAmount float64, tokenAddress string, tokenAmount float64, contractHash string) (common.IssuanceResponse, error) {
	return assetIssuance(ctx, chain, DefaultConfidentialIssuance, issuerID, common.IssuanceRequest{
		Mode:               common.AssetIssuanceModeWithTokenWithContract,
		AssetPublicAddress: assetAddress,
		AssetIssuedAmount:  assetAmount,
		TokenPublicAddress: tokenAddress,
		TokenIssuedAmount:  tokenAmount,
		ContractHash:       contractHash,
	})
}

func assetIssuance(ctx context.Context, chain string, confidential bool, issuerID uint64, request common.IssuanceRequest) (common.IssuanceResponse, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.client.assetIssuance")

	request.Chain = chain
	request.IssuerID = issuerID

	if !request.IsValid() {
		return common.IssuanceResponse{}, errors.New("Invalid Issuance info")
	}

	var result common.IssuanceResponse
	err := messaging.RequestMessage(ctx, appcontext.AppName(ctx), common.AssetIssuanceSubject, &request, &result)
	if err != nil {
		log.WithError(err).
			Error("RequestMessage failed")
		return common.IssuanceResponse{}, messaging.ErrRequestFailed
	}

	return result, nil
}
