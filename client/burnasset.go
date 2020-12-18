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

// AssetBurn burns amount of asset by spending it to an unspendable output
func AssetBurn(ctx context.Context, chain string, issuerID uint64, assetToBurn string, amountToBurn float64) (common.BurnResponse, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.client.AssetBurn")

	var request common.BurnRequest

	request.Chain = chain
	request.IssuerID = issuerID

	request.Asset = assetToBurn
	request.Amount = amountToBurn

	var result common.BurnResponse
	err := messaging.RequestMessage(ctx, appcontext.AppName(ctx), common.AssetBurnSubject, &request, &result)
	if err != nil {
		log.WithError(err).
			Error("RequestMessage failed")
		return common.BurnResponse{}, messaging.ErrRequestFailed
	}

	return result, nil
}
