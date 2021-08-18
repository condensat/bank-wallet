// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"

	"github.com/condensat/bank-wallet/common"
	"github.com/condensat/bank-wallet/ln"
)

func main() {
	ctx := context.Background()
	args := parseArgs(ctx)

	client := createLightningClient(ctx, args.Common)
	if client == nil {
		log.Fatalf("Failed to create client")
	}

	ctx = context.WithValue(ctx, "LightningClientKey", client)

	Run(ctx, args)
}

func Run(ctx context.Context, args Args) {
	var err error
	switch args.Command {
	case GetInfo:
		err = getInfo(ctx)

	case KeySend:
		err = keySend(ctx, args.KeySend)

	case Pay:
		err = pay(ctx, args.Pay)

	case DecodePay:
		err = deocdepay(ctx, args.DecodePay)

	case ListInvoices:
		err = listInvoices(ctx, args.ListInvoices)

	default:
		printUsage(1)
	}

	if err != nil {
		log.Fatalf("Error while processing command. (%s)\n", err)
	}
}

func createLightningClient(ctx context.Context, args CommonArg) common.LightningClient {
	return ln.NewWithTorEndpoint(ctx,
		args.torProxy,
		args.endpoint,
		args.macaroon,
	)
}
