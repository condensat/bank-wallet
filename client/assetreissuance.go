// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package client

import (
	"context"

	"github.com/condensat/bank-core/appcontext"
	"github.com/condensat/bank-core/logger"
	"github.com/condensat/bank-core/messaging"
	"github.com/condensat/bank-wallet/common"
)

// AssetReissuance reissues an asset if provided with a token input
func AssetReissuance(ctx context.Context, chain string, issuerID uint64, assetID, assetAddress string, assetAmount float64) (common.ReissuanceResponse, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.client.assetReissuance")

	var request common.ReissuanceRequest

	request.Chain = chain
	request.IssuerID = issuerID

	request.AssetID = assetID
	request.AssetPublicAddress = assetAddress
	request.AssetIssuedAmount = assetAmount

	var result common.ReissuanceResponse
	err := messaging.RequestMessage(ctx, appcontext.AppName(ctx), common.AssetReissuanceSubject, &request, &result)
	if err != nil {
		log.WithError(err).
			Error("RequestMessage failed")
		return common.ReissuanceResponse{}, messaging.ErrRequestFailed
	}

	return result, nil
}
