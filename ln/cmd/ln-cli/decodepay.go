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
	DecodePay = Command("decodePay")
)

type DecodePayArg struct {
	invoice string
}

func decodePayArgs(args *DecodePayArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("decodePay", flag.ExitOnError)

	cmd.StringVar(&args.invoice, "invoice", "", "Invoice to decode (Bolt	11)")

	return cmd
}

func deocdepay(ctx context.Context, args DecodePayArg) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	res, err := client.DecodePay(ctx, args.invoice)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&res)
	fmt.Printf("%s\n", obj)

	return err
}
