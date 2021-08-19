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
	GenInvoice = Command("genInvoice")
)

type GenInvoiceArg struct {
	amount      int
	label       string
	description string
	expiry      int
	private     bool
}

func genInvoiceArgs(args *GenInvoiceArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("genInvoice", flag.ExitOnError)

	cmd.IntVar(&args.amount, "amount", 0, "Invoice Amount (msats)")
	cmd.StringVar(&args.label, "label", "", "Invoice label")
	cmd.StringVar(&args.description, "description", "", "Invoice description")

	cmd.IntVar(&args.expiry, "expiry", 0, "Invoice expiry (seconds)")
	cmd.BoolVar(&args.private, "private", true, "Invoice private")

	return cmd
}

func genInvoice(ctx context.Context, args GenInvoiceArg) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	res, err := client.GenInvoice(ctx, args.amount, args.label, args.description, args.expiry, args.private)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&res)
	fmt.Printf("%s\n", obj)

	return err
}
