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
	GetInfo = Command("getInfo")
)

type GetInfoArg struct {
}

func getInfoArgs(args *GetInfoArg) *flag.FlagSet {
	cmd := flag.NewFlagSet("getInfo", flag.ExitOnError)

	return cmd
}

func getInfo(ctx context.Context) error {
	client := ctx.Value("LightningClientKey").(common.LightningClient)

	info, err := client.GetInfo(ctx)
	if err != nil {
		return err
	}

	obj, _ := json.Marshal(&info)
	fmt.Printf("%s\n", obj)

	return err
}
