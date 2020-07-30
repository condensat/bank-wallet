// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package chain

import (
	"context"
	"errors"
	"sort"

	"github.com/condensat/bank-core/cache"
	"github.com/condensat/bank-core/logger"

	"github.com/condensat/bank-wallet/common"

	"github.com/condensat/bank-core/utils"

	"github.com/sirupsen/logrus"
)

const (
	AddressBatchSize = 16 // maximum address count for RPC requests
)

var (
	ErrChainClientNotFound = errors.New("ChainClient Not Found")
)

type ChainState struct {
	Chain  string
	Height uint64
}

type TransactionInfo struct {
	TxID          string
	Vout          int64
	Asset         string
	Amount        float64
	Confirmations int64
}

type AddressInfo struct {
	Chain         string
	PublicAddress string
	Mined         uint64 // 0 unknown, 1 mempool, BlockHeight
	Transactions  []TransactionInfo
}

func GetNewAddress(ctx context.Context, chain, account string) (string, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.GetNewAddress")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return "", ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return "", cache.ErrLockError
	}
	defer lock.Unlock()

	return client.GetNewAddress(ctx, account)
}

func GetAddressInfo(ctx context.Context, chain, address string) (common.AddressInfo, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.GetAddressInfo")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return common.AddressInfo{}, ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return common.AddressInfo{}, cache.ErrLockError
	}
	defer lock.Unlock()

	info, err := client.GetAddressInfo(ctx, address)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return common.AddressInfo{}, cache.ErrLockError
	}

	info.Chain = chain

	return info, nil
}

func LockUnspent(ctx context.Context, chain string, unlock bool, utxos ...common.TransactionInfo) error {
	log := logger.Logger(ctx).WithField("Method", "wallet.LockUnspent")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return cache.ErrLockError
	}
	defer lock.Unlock()

	err = client.LockUnspent(ctx, unlock, utxos...)
	if err != nil {
		log.WithError(err).
			Error("Failed to LockUnspent")
		return cache.ErrLockError
	}

	return nil
}

func ListLockUnspent(ctx context.Context, chain string) ([]common.TransactionInfo, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.ListLockUnspent")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return nil, ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return nil, cache.ErrLockError
	}
	defer lock.Unlock()

	list, err := client.ListLockUnspent(ctx)
	if err != nil {
		log.WithError(err).
			Error("Failed to ListLockUnspent")
		return nil, err
	}

	return list, nil
}

func GetTransaction(ctx context.Context, chain string, txID string) (common.TransactionInfo, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.GetTransaction")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return common.TransactionInfo{}, ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return common.TransactionInfo{}, cache.ErrLockError
	}
	defer lock.Unlock()

	result, err := client.GetTransaction(ctx, txID)
	if err != nil {
		log.WithError(err).
			Error("Failed to GetTransaction")
		return common.TransactionInfo{}, cache.ErrLockError
	}

	return result, nil
}

func FetchChainsState(ctx context.Context, chains ...string) ([]ChainState, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.FetchChainsState")

	var result []ChainState
	for _, chain := range chains {
		state, err := fetchChainState(ctx, chain)
		if err != nil {
			continue
		}

		result = append(result, state)
	}

	log.
		WithField("Count", len(result)).
		Debug("Chains State Fetched")

	return result, nil
}

func fetchChainState(ctx context.Context, chain string) (ChainState, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.fetchChainState")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return ChainState{}, ErrChainClientNotFound
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return ChainState{}, cache.ErrLockError
	}
	defer lock.Unlock()

	blockCount, err := client.GetBlockCount(ctx)
	if err != nil {
		return ChainState{}, err
	}

	log.
		WithFields(logrus.Fields{
			"Chain":  chain,
			"Height": blockCount,
		}).Info("Client RPC")

	return ChainState{
		Chain:  chain,
		Height: uint64(blockCount),
	}, nil
}

func FetchChainAddressesInfo(ctx context.Context, state ChainState, minConf, maxConf uint64, publicAddresses ...string) ([]AddressInfo, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.FetchChainAddresses")

	log = log.WithFields(logrus.Fields{
		"Chain":  state.Chain,
		"Height": state.Height,
	})

	client := common.ChainClientFromContext(ctx, state.Chain)
	if client == nil {
		return nil, ErrChainClientNotFound
	}

	if len(publicAddresses) == 0 {
		log.Debug("No addresses provided")
		return nil, nil
	}

	// Acquire Lock
	lock, err := cache.LockChain(ctx, state.Chain)
	if err != nil {
		log.WithError(err).
			Error("Failed to lock chain")
		return nil, cache.ErrLockError
	}
	defer lock.Unlock()

	if minConf > maxConf {
		maxConf, minConf = minConf, maxConf
	}

	firsts := make(map[string]*AddressInfo)
	batches := utils.BatchString(AddressBatchSize, publicAddresses...)
	for _, batch := range batches {

		// RPC requets
		list, err := client.ListUnspent(ctx, int(minConf), int(maxConf), batch...)
		if err != nil {
			log.WithError(err).
				Error("Failed to ListUnspent")
			return nil, err
		}

		lockedUtxos, err := client.ListLockUnspent(ctx)
		if err != nil {
			log.WithError(err).
				Error("Failed to ListLockUnspent")
			return nil, err
		}
		for _, utxo := range lockedUtxos {
			tx, err := client.GetTransaction(ctx, utxo.TxID)
			if err != nil {
				log.WithError(err).
					Error("Failed to GetTransaction")
				return nil, err
			}
			list = append(list, tx)
		}

		// Order oldest first
		sort.Slice(list, func(i, j int) bool {
			return list[i].Confirmations > list[j].Confirmations
		})

		for _, utxo := range list {
			// create if address is already not exists
			if _, ok := firsts[utxo.Address]; !ok {

				// zero confirmation mean in mempool
				var blockHeight uint64
				if utxo.Confirmations > 0 {
					blockHeight = state.Height - uint64(utxo.Confirmations)
				}

				// create new map entry
				firsts[utxo.Address] = &AddressInfo{
					Chain:         state.Chain,
					PublicAddress: utxo.Address,
					Mined:         blockHeight,
				}
			}

			// append TxID to transactions
			addr := firsts[utxo.Address]
			addr.Transactions = append(addr.Transactions, TransactionInfo{
				TxID:          utxo.TxID,
				Vout:          utxo.Vout,
				Asset:         utxo.Asset,
				Amount:        utxo.Amount,
				Confirmations: utxo.Confirmations,
			})
		}
	}

	var result []AddressInfo
	for _, utxo := range firsts {
		if utxo == nil {
			continue
		}
		result = append(result, *utxo)
	}

	return result, nil
}

func SpendFunds(ctx context.Context, chain string, spendInfos []common.SpendInfo) (common.SpendTx, error) {
	log := logger.Logger(ctx).WithField("Method", "wallet.SpendFunds")

	log = log.WithField("Chain", chain)

	client := common.ChainClientFromContext(ctx, chain)
	if client == nil {
		return common.SpendTx{}, ErrChainClientNotFound
	}

	// Create, Fund, Sign & Broadcast transaction
	tx, err := client.SpendFunds(ctx, nil, spendInfos)
	if err != nil {
		log.WithError(err).
			Error("Failed to SpendFunds")
		return common.SpendTx{}, err
	}

	return tx, nil
}
