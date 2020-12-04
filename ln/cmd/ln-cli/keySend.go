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
	KeySend = Command("keySend")
)

type KeySendArg struct {
	pubKey string
	amount int
	label  string
}

func keySendArgs(args *KeySendArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("keySend", flag.ExitOnError)

	cmd.StringVar(&args.pubKey, "pubKey", "", "KeySend pubKey")
	cmd.IntVar(&args.amount, "amount", 0, "KeySend amount (sat)")
	cmd.StringVar(&args.label, "label", "", "KeySend label")

	return cmd
}

func keySend(ctx context.Context, args KeySendArg) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	res, err := client.KeySend(ctx,
		args.pubKey,
		args.amount*1000,
		args.label,
	)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&res)
	fmt.Printf("%s\n", obj)

	return err
}
