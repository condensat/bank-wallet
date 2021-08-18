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
	ListInvoices = Command("listInvoices")
)

type ListInvoicesArg struct {
	label string
}

func listInvoicesArgs(args *ListInvoicesArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("listInvoices", flag.ExitOnError)

	cmd.StringVar(&args.label, "label", "", "Filter Invoice by label")

	return cmd
}

func listInvoices(ctx context.Context, args ListInvoicesArg) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	res, err := client.ListInvoices(ctx, args.label)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&res)
	fmt.Printf("%s\n", obj)

	return err
}
