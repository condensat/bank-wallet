// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/condensat/bank-wallet/common"
	"github.com/condensat/bank-wallet/rpc"

	"github.com/condensat/bank-wallet/bitcoin/commands"

	dotenv "github.com/joho/godotenv"
)

func init() {
	_ = dotenv.Load()
}

type ChainOption struct {
	Chain    string
	HostName string
	Port     int
	User     string
	Pass     string
}

type ChainsOptions struct {
	Chains []ChainOption
}

func main() {
	var command string

	flag.StringVar(&command, "command", "", "Sub command to start")

	flag.Parse()

	ctx := context.Background()

	// add CryptoMode to context
	ctx = common.CryptoModeContext(ctx, common.CryptoModeBitcoinCore)

	var err error
	switch command {

	// Bitcoin Standard

	case "RawTransactionBitcoin":
		err = RawTransactionBitcoin(ctx)

	// Liquid Elements

	case "RawTransactionElements":
		err = RawTransactionElements(ctx)

	default:
		log.Fatalf("Unknown command %s.", command)
	}

	if err != nil {
		log.Fatalf("Failed to process command. %v", err)
	}
}

// Bitcoin Standard

func RawTransactionBitcoin(ctx context.Context) error {
	rpcClient := bitcoinRpcClient()

	destAddress := "bcrt1qjlw9gfrqk0w2ljegl7vwzrt2rk7sst8d4hm7n9"

	hex, err := commands.CreateRawTransaction(ctx, rpcClient, nil, []commands.SpendInfo{
		{Address: destAddress, Amount: 0.003},
	}, nil)
	if err != nil {
		return err
	}
	log.Printf("CreateRawTransaction: %s\n", hex)

	rawTx, err := commands.DecodeRawTransaction(ctx, rpcClient, hex)
	if err != nil {
		return err
	}
	decoded, err := commands.ConvertToRawTransactionBitcoin(rawTx)
	if err != nil {
		return err
	}
	log.Printf("DecodeRawTransaction: %+v\n", decoded)

	funded, err := commands.FundRawTransaction(ctx, rpcClient, hex)
	if err != nil {
		return err
	}
	log.Printf("FundRawTransaction: %+v\n", funded)

	rawTx, err = commands.DecodeRawTransaction(ctx, rpcClient, commands.Transaction(funded.Hex))
	if err != nil {
		return err
	}
	decoded, err = commands.ConvertToRawTransactionBitcoin(rawTx)
	if err != nil {
		return err
	}
	log.Printf("FundRawTransaction Hex: %+v\n", decoded)

	addressMap := make(map[commands.Address]commands.Address)
	for _, in := range decoded.Vin {

		txInfo, err := commands.GetTransaction(ctx, rpcClient, in.Txid, true)
		if err != nil {
			return err
		}

		addressMap[txInfo.Address] = txInfo.Address
		for _, d := range txInfo.Details {
			address := commands.Address(d.Address)
			addressMap[address] = address
		}
	}

	signed, err := commands.SignRawTransactionWithWallet(ctx, rpcClient, commands.Transaction(funded.Hex))
	if err != nil {
		return err
	}
	if !signed.Complete {
		return errors.New("SignRawTransactionWithWallet failed")
	}
	log.Printf("Signed transaction is: %+v\n", signed.Hex)

	accepted, err := commands.TestMempoolAccept(ctx, rpcClient, signed.Hex)
	if err != nil {
		return err
	}

	log.Printf("Accepted in the mempool: %+v\n", accepted.Allowed)
	if !accepted.Allowed {
		log.Printf("Reject-reason: %+v", accepted.Reason)
		return errors.New("TestMempoolAccept failed")
	}

	return nil
}

// Liquid Elements

func RawTransactionElements(ctx context.Context) error {
	rpcClient := elementsRpcClient()

	destAddress := "el1qqw8rsv0nxl3mvgztna2n6fyz37wnucyt9yv5qcwty0af6e3yfaj5ke0hnadd96tp03nz8tltz4yxq39yqal9jjq9ry25gjhpw"
	changeAddress := "el1qqdhtdknl5wazd2jysqwhun7tyx8zycygvtdyz0hg9tnr96m00ateqrewrzncus3hwdfvj9t9ehf45k5y700pjsdfc44khklma"
	// We create 2 LBTC outputs, which might be a bit unnecessary
	hex, err := commands.CreateRawTransaction(ctx, rpcClient, nil, []commands.SpendInfo{
		{Address: destAddress, Amount: 0.001},
	}, nil)
	if err != nil {
		return err
	}
	log.Printf("CreateRawTransaction: %s\n", hex)

	rawTx, err := commands.DecodeRawTransaction(ctx, rpcClient, hex)
	if err != nil {
		return err
	}
	decoded, err := commands.ConvertToRawTransactionLiquid(rawTx)
	if err != nil {
		return err
	}
	log.Printf("DecodeRawTransaction: %+v\n", decoded)

	funded, err := commands.FundRawTransactionWithOptions(ctx,
		rpcClient,
		hex,
		commands.FundRawTransactionOptions{
			ChangeAddress:   changeAddress,
			IncludeWatching: true,
		},
	)
	if err != nil {
		return err
	}
	log.Printf("FundRawTransaction: %+v\n", funded)

	blinded, err := commands.BlindRawTransaction(ctx, rpcClient, commands.Transaction(funded.Hex))
	if err != nil {
		return err
	}

	log.Printf("Blinded transaction OK\n")

	signed, err := commands.SignRawTransactionWithWallet(ctx, rpcClient, commands.Transaction(blinded))
	if err != nil {
		return err
	}
	if !signed.Complete {
		return errors.New("SignRawTransactionWithWallet failed")
	}

	accepted, err := commands.TestMempoolAccept(ctx, rpcClient, signed.Hex)
	if err != nil {
		return err
	}

	log.Printf("Accepted in the mempool: %+v\n", accepted.Allowed)
	if !accepted.Allowed {
		log.Printf("Reject-reason: %+v", accepted.Reason)
		return errors.New("TestMempoolAccept failed")
	}

	return nil
}

// Helpers

func bitcoinRpcClient() commands.RpcClient {
	hostname := os.Getenv("BITCOIN_TESTNET_HOSTNAME")
	if len(hostname) == 0 {
		hostname = "bitcoin_testnet"
	}
	port, _ := strconv.Atoi(os.Getenv("BITCOIN_TESTNET_PORT"))
	if port == 0 {
		port = 18332
	}
	user := os.Getenv("BITCOIN_TESTNET_USER")
	if len(user) == 0 {
		user = "bank-wallet"
	}
	password := os.Getenv("BITCOIN_TESTNET_PASSWORD")
	if len(password) == 0 {
		password = "password1"
	}

	return createRpcClient(hostname, port, user, password)
}

func elementsRpcClient() commands.RpcClient {
	hostname := os.Getenv("ELEMENTS_REGTEST_HOSTNAME")
	if len(hostname) == 0 {
		hostname = "elements_regtest"
	}
	port, _ := strconv.Atoi(os.Getenv("ELEMENTS_REGTEST_PORT"))
	if port == 0 {
		port = 28432
	}
	user := os.Getenv("ELEMENTS_REGTEST_USER")
	if len(user) == 0 {
		user = "bank-wallet"
	}
	password := os.Getenv("ELEMENTS_REGTEST_PASSWORD")
	if len(password) == 0 {
		password = "password1"
	}

	return createRpcClient(hostname, port, user, password)
}

func createRpcClient(hostname string, port int, user, password string) commands.RpcClient {
	rpcClient := rpc.New(rpc.Options{
		ServerOptions: common.ServerOptions{Protocol: "http", HostName: hostname, Port: port},
		User:          user,
		Password:      password,
	}).Client

	_, err := commands.GetBlockCount(context.Background(), rpcClient)
	if err != nil {
		log.Fatalf("Rpc call failed. %s.", err)
	}

	return rpcClient
}
