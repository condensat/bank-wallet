// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/condensat/bank-wallet/common"
)

const (
	Pay = Command("pay")
)

type PayArg struct {
	invoice string
}

func payArgs(args *PayArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("pay", flag.ExitOnError)

	cmd.StringVar(&args.invoice, "invoice", "", "Pay invoice")

	return cmd
}

func pay(ctx context.Context, args PayArg) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	res, err := client.Pay(ctx, args.invoice)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&res)
	fmt.Printf("%s\n", obj)

	return err
}
