// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"time"

	"github.com/condensat/bank-core/appcontext"
	"github.com/condensat/bank-core/cache"
	"github.com/condensat/bank-core/logger"
	"github.com/condensat/bank-core/messaging"
	"github.com/condensat/bank-core/monitor/processus"

	"github.com/condensat/bank-wallet/client"

	"github.com/sirupsen/logrus"
)

type Args struct {
	App appcontext.Options

	Redis cache.RedisOptions
	Nats  messaging.NatsOptions
}

func parseArgs() Args {
	var args Args
	appcontext.OptionArgs(&args.App, "BankWalletCli")

	cache.OptionArgs(&args.Redis)
	messaging.OptionArgs(&args.Nats)

	flag.Parse()

	return args
}

func WalletCli(ctx context.Context) {
	log := logger.Logger(ctx).WithField("Method", "WalletCli")

	// list all currencies
	addr, err := client.CryptoAddressNextDeposit(ctx, "bitcoin-mainnet", 42)
	if err != nil {
		panic(err)
	}

	log.WithFields(logrus.Fields{
		"Chain":         addr.Chain,
		"AccountID":     addr.AccountID,
		"PublicAddress": addr.PublicAddress,
	}).Info("CryptoAddress NextDeposit")
}

func main() {
	args := parseArgs()

	ctx := context.Background()
	ctx = appcontext.WithOptions(ctx, args.App)
	ctx = appcontext.WithCache(ctx, cache.NewRedis(ctx, args.Redis))
	ctx = appcontext.WithWriter(ctx, logger.NewRedisLogger(ctx))
	ctx = appcontext.WithMessaging(ctx, messaging.NewNats(ctx, args.Nats))
	ctx = appcontext.WithProcessusGrabber(ctx, processus.NewGrabber(ctx, 15*time.Second))

	WalletCli(ctx)
}