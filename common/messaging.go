// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package common

const (
	chanPrefix = "Condensat.Wallet."

	CryptoAddressNextDepositSubject = chanPrefix + "CryptoAddress.NextDeposit"
	CryptoAddressNewDepositSubject  = chanPrefix + "CryptoAddress.NewDeposit"
	AddressInfoSubject              = chanPrefix + "CryptoAddress.AddressInfo"

	WalletStatusSubject = chanPrefix + "WalletStatus"
	WalletListSubject   = chanPrefix + "WalletList"

	AssetListIssuancesSubject = chanPrefix + "Asset.ListIssuances"
	AssetIssuanceSubject      = chanPrefix + "Asset.Issuance"
	AssetReissuanceSubject    = chanPrefix + "Asset.Reissuance"
	AssetBurnSubject          = chanPrefix + "Asset.Burn"
)
