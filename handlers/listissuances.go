// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package handlers

import (
	"context"
	"errors"

	"github.com/condensat/bank-core/appcontext"
	"github.com/condensat/bank-core/logger"

	"github.com/condensat/bank-wallet/common"

	"github.com/condensat/bank-core/cache"
	"github.com/condensat/bank-core/messaging"

	"github.com/sirupsen/logrus"
)

func ListIssuances(ctx context.Context, request common.ListIssuancesRequest) ([]common.IssuanceInfo, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.AssetIssuance")

	chainHandler := ChainHandlerFromContext(ctx)
	if chainHandler == nil {
		log.Error("Failed to ChainHandlerFromContext")
		return []common.IssuanceInfo{}, errors.New("Something's wrong with the chainHandler")
	}

	return chainHandler.ListIssuances(ctx, request)
}

func OnListIssuances(ctx context.Context, subject string, message *messaging.Message) (*messaging.Message, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.OnListIssuances")
	log = log.WithFields(logrus.Fields{
		"Subject": subject,
	})

	var request common.ListIssuancesRequest
	return messaging.HandleRequest(ctx, appcontext.AppName(ctx), message, &request,
		func(ctx context.Context, _ messaging.BankObject) (messaging.BankObject, error) {
			log = log.WithFields(logrus.Fields{
				"Chain":    request.Chain,
				"IssuerID": request.IssuerID,
			})

			list, err := ListIssuances(ctx, request)
			if err != nil {
				log.WithError(err).
					Errorf("Failed to ListIssuances")
				return nil, cache.ErrInternalError
			}

			return &common.IssuanceList{
				Chain:     request.Chain,
				IssuerID:  request.IssuerID,
				Issuances: list,
			}, nil
		})
}
