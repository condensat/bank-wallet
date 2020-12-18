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

	"github.com/sirupsen/logrus"
)

func ListIssuances(ctx context.Context, chain string, issuerID uint64, asset string) (common.IssuanceList, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.client.ListIssuances")

	var request common.ListIssuancesRequest
	var response common.IssuanceList

	request.Chain = chain
	request.IssuerID = issuerID
	request.Asset = asset
	err := messaging.RequestMessage(ctx, appcontext.AppName(ctx), common.AssetListIssuancesSubject, &request, &response)
	if err != nil {
		log.WithError(err).
			Error("RequestMessage failed")
		return common.IssuanceList{}, messaging.ErrRequestFailed
	}

	log.WithFields(logrus.Fields{
		"Issuer ID": response.IssuerID,
		"Count":     len(response.Issuances),
	}).Debug("Issuances info")

	return response, nil
}
