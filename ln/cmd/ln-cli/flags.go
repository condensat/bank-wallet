// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	dotenv "github.com/joho/godotenv"
)

func init() {
	_ = dotenv.Load()
}

func printUsage(code int) {
	fmt.Println("Use command [getInfo, keySend]")
	os.Exit(code)
}

type Command string

type CommonArg struct {
	torProxy string
	endpoint string
	macaroon string
}

type Args struct {
	Command Command
	Common  CommonArg

	GetInfo GetInfoArg
	KeySend KeySendArg
}

func commonArgs(cmd *flag.FlagSet, args *CommonArg) {
	cmd.StringVar(&args.torProxy, "torproxy", "", "Tor Proxy for hidden services access")
	cmd.StringVar(&args.endpoint, "endpoint", "", "Lightning node onion address")
	cmd.StringVar(&args.macaroon, "macaroon", "", "macaroon in hex")
}

func parseArgs(ctx context.Context) Args {
	var args Args

	if len(os.Args) == 1 {
		printUsage(1)
	}
	args.Command = Command(os.Args[1])

	var cmd *flag.FlagSet
	switch args.Command {
	case GetInfo:
		cmd = getInfoArgs(&args.GetInfo)
	case KeySend:
		cmd = keySendArgs(&args.KeySend)

	default:
		printUsage(2)
	}

	commonArgs(cmd, &args.Common)

	err := cmd.Parse(os.Args[2:])
	if err != nil {
		printUsage(3)
	}

	// Load from .env
	fromStringEnv("LN_CLI_TORPROXY", &args.Common.torProxy)
	fromStringEnv("LN_CLI_ENDPOINT", &args.Common.endpoint)
	fromStringEnv("LN_CLI_MACAROON", &args.Common.macaroon)

	return args
}

func fromStringEnv(key string, value *string) {
	if len(*value) != 0 {
		return
	}
	*value = os.Getenv(key)
}
