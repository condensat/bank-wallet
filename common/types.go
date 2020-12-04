// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package common

import (
	"github.com/condensat/bank-core/messaging"
)

type CryptoMode string
type AssetIssuanceMode string

const (
	CryptoModeBitcoinCore CryptoMode = "bitcoin-core"
	CryptoModeCryptoSsm   CryptoMode = "crypto-ssm"
)

type ServerOptions struct {
	Protocol string
	HostName string
	Port     int
}

const (
	AssetIssuanceModeWithAsset             AssetIssuanceMode = "asset-only"
	AssetIssuanceModeWithToken             AssetIssuanceMode = "with-token"
	AssetIssuanceModeWithContract          AssetIssuanceMode = "with-contract"
	AssetIssuanceModeWithTokenWithContract AssetIssuanceMode = "with-token-and-contract"
)

const (
	ElementsRegtestHash string = "b2e15d0d7a0c94e4e2ce0fe6e8691b9e451377f6e46e8045a86f7c4b5d4f0f23"
)

type CryptoAddress struct {
	CryptoAddressID  uint64
	Chain            string
	AccountID        uint64
	PublicAddress    string
	Unconfidential   string
	IgnoreAccounting bool
}

type SsmAddress struct {
	Chain       string
	Address     string
	PubKey      string
	BlindingKey string
}

type ElementsBlindingInfo struct {
	AssetBlinder string
}

type TransactionInfo struct {
	Chain         string
	Account       string
	Address       string
	Asset         string
	TxID          string
	Vout          int64
	Amount        float64
	Confirmations int64
	Spendable     bool
	Blinding      ElementsBlindingInfo
}

type AddressInfo struct {
	Chain          string
	PublicAddress  string
	Unconfidential string
	IsValid        bool
}

type UTXOInfo struct {
	TxID   string
	Vout   int
	Asset  string
	Amount float64
	Locked bool
}

type SpendAssetInfo struct {
	Hash          string
	ChangeAddress string
	ChangeAmount  float64
}

type SpendInfo struct {
	PublicAddress string
	Amount        float64
	// Asset optional
	Asset SpendAssetInfo
}

type SpendTx struct {
	TxID string
}

type ListIssuancesRequest struct {
	Chain    string
	IssuerID uint64
	Asset    string
}

type IssuanceInfo struct {
	TxID         string  // TxID of the issuance
	Entropy      string  // We need it for reissuances
	Asset        string  // Asset ID (64B hex)
	Token        string  // reissuance token ID, computed from asset ID
	Vin          int     // index of the input the issuance is "hooked" to
	AssetAmount  float64 // amount issued of the asset
	TokenAmount  float64 // amount of reissuance token (can't be reissued)
	IsReissuance bool    // false == initial issuance
	AssetBlinds  string  // blinding factor for the amount of asset issued
	TokenBlinds  string  // blinding factor for the amount of token issued
}

type IssuanceList struct {
	Chain     string
	IssuerID  uint64
	Issuances []IssuanceInfo
}

type IssuanceRequest struct {
	Chain              string            // mainly elements-regtest or LiquidV1 now, but can be useful for other chains later
	IssuerID           uint64            // User ID used for communication with our db
	Mode               AssetIssuanceMode // Issue an asset either with a reissuance token, a contract hash or both
	BlindIssuance      bool              // Issuance can be blinded or not
	AssetPublicAddress string            // Address we send the newly issued asset to
	AssetIssuedAmount  float64           // Max 21_000_000.0, but can be reissued many times

	// Optional
	TokenPublicAddress string  // Address we send the reissuance token to
	TokenIssuedAmount  float64 // I'd recommend it to be either 0 or 0.00000001 (1 sat)
	ContractHash       string  // 32B hash we can commit directly inside the asset ID
}

type IssuanceResponse struct {
	Chain     string   // mainly elements-regtest or LiquidV1 now, but can be useful for other chains later
	IssuerID  uint64   // User ID used for communication with our db
	AssetID   string   // This is the hex 64B Identifier of the asset. It is computed determinastically from a txid, a vout and an optional contract hash
	TokenID   string   // hex 64B identifier of the token that allows to reissue the asset
	TxID      string   // ID of the issuance transaction
	Vin       UTXOInfo // Txid and vout of the input the issuance is hooked to, used to compute asset ID with contract hash if any
	AssetVout int      // The vout of the new asset
	TokenVout int      // The vout of the token. We need this for reissuance
	Entropy   string   // Entropy is calculated with the issuance vin and contract hash if any
}

type ReissuanceRequest struct {
	Chain              string
	IssuerID           uint64  // User ID used for communication with our db
	AssetID            string  // Asset that has been reissued
	AssetPublicAddress string  // Address to reissue assets to
	AssetIssuedAmount  float64 // Max 21_000_000.0, but can be reissued many times

	TokenID            string  // hex 64B identifier of the token that allows to reissue the asset
	TokenPublicAddress string  // Address to send tokens to
	TokenAmount        float64 // amount locked in the spent UTXO

	Entropy      string // some entropy determined at issuance, and that we can get from listissuance() call
	AssetBlinder string // From the output being spent
	InputIndex   int    // To which Vin we want to attach the reissuance
}

type ReissuanceResponse struct {
	Chain     string
	IssuerID  uint64
	TxID      string // txid of the reissuance transaction
	AssetVout int    // vout of the newly issued assets
	TokenVout int    // vout of the reissuance token: this is what we will need next time to issue
}

type BurnRequest struct {
	Chain    string
	IssuerID uint64
	Asset    string  // asset to burn
	Amount   float64 // amount of asset to burn
}

type BurnResponse struct {
	Chain    string
	IssuerID uint64
	TxID     string // txid of the transaction that burned the asset
	Vout     int    // index of the OP_RETURN output that burn the asset
}

type WalletInfo struct {
	Chain  string
	Height int
	UTXOs  []UTXOInfo
}

type WalletStatus struct {
	Wallets []WalletInfo
}

func (p *CryptoAddress) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *CryptoAddress) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *AddressInfo) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *AddressInfo) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *WalletInfo) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *WalletInfo) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *WalletStatus) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *WalletStatus) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *ListIssuancesRequest) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *ListIssuancesRequest) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *IssuanceList) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *IssuanceList) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *IssuanceRequest) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *IssuanceRequest) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *IssuanceResponse) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *IssuanceResponse) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *ReissuanceRequest) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *ReissuanceRequest) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *ReissuanceResponse) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *ReissuanceResponse) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *BurnRequest) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *BurnRequest) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *BurnResponse) Encode() ([]byte, error) {
	return messaging.EncodeObject(p)
}

func (p *BurnResponse) Decode(data []byte) error {
	return messaging.DecodeObject(data, messaging.BankObject(p))
}

func (p *IssuanceRequest) IsValid() bool {
	switch p.Mode {
	case AssetIssuanceModeWithAsset:
		if len(p.AssetPublicAddress) == 0 {
			return false
		}
		if p.AssetIssuedAmount <= 0.0 {
			return false
		}
		if len(p.TokenPublicAddress) != 0 {
			return false
		}
		if p.TokenIssuedAmount > 0.0 {
			return false
		}
		if len(p.ContractHash) != 0 {
			return false
		}
		return true
	case AssetIssuanceModeWithToken:
		if len(p.AssetPublicAddress) == 0 {
			return false
		}
		if p.AssetIssuedAmount <= 0.0 {
			return false
		}
		if len(p.TokenPublicAddress) == 0 {
			return false
		}
		if p.TokenIssuedAmount <= 0.0 {
			return false
		}
		if len(p.ContractHash) != 0 {
			return false
		}
		return true
	case AssetIssuanceModeWithContract:
		if len(p.AssetPublicAddress) == 0 {
			return false
		}
		if p.AssetIssuedAmount <= 0.0 {
			return false
		}
		if len(p.TokenPublicAddress) != 0 {
			return false
		}
		if p.TokenIssuedAmount > 0.0 {
			return false
		}
		if len(p.ContractHash) == 0 {
			return false
		}
		return true
	case AssetIssuanceModeWithTokenWithContract:
		if len(p.AssetPublicAddress) == 0 {
			return false
		}
		if p.AssetIssuedAmount <= 0.0 {
			return false
		}
		if len(p.TokenPublicAddress) == 0 {
			return false
		}
		if p.TokenIssuedAmount <= 0.0 {
			return false
		}
		if len(p.ContractHash) == 0 {
			return false
		}
		return true
	default:
		return false
	}
}

// Lightning

type ResponseError struct {
	Code     int    `json:"code,omitempty"`
	FullType string `json:"fullType,omitempty"`
	Message  string `json:"message,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Attempts []struct {
		Status       string `json:"status,omitempty"`
		Failreason   string `json:"failreason,omitempty"`
		Partid       int    `json:"partid,omitempty"`
		Amount       string `json:"amount,omitempty"`
		ParentPartid int    `json:"parent_partid,omitempty"`
	} `json:"attempts,omitempty"`
}

type GetInfoResponse struct {
	ID                  string `json:"id"`
	Alias               string `json:"alias"`
	Color               string `json:"color"`
	NumPeers            int    `json:"num_peers"`
	NumPendingChannels  int    `json:"num_pending_channels"`
	NumActiveChannels   int    `json:"num_active_channels"`
	NumInactiveChannels int    `json:"num_inactive_channels"`
	Address             []struct {
		Type    string `json:"type"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"address"`
	Binding []struct {
		Type    string `json:"type"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"binding"`
	Version               string `json:"version"`
	Blockheight           int    `json:"blockheight"`
	Network               string `json:"network"`
	MsatoshiFeesCollected int    `json:"msatoshi_fees_collected"`
	FeesCollectedMsat     string `json:"fees_collected_msat"`
	LightningDir          string `json:"lightning-dir"`
	APIVersion            string `json:"api_version"`

	ResponseError
	Error *ResponseError `json:"error,omitempty"`
}

type KeySendArgs struct {
	PubKey string `json:"pubkey"`
	Amount int    `json:"amount"`
	Label  string `json:"label"`

	MaxFeePercent float64 `json:"maxfeepercent,omitempty"`
	RetryFor      int     `json:"retry_for,omitempty"`
	MaxDelay      int     `json:"maxdelay,omitempty"`
	ExemptFee     int     `json:"exemptfee,omitempty"`
}

type KeySendResponse struct {
	AmountMsat      string  `json:"amount_msat,omitempty"`
	AmountSentMsat  string  `json:"amount_sent_msat,omitempty"`
	CreatedAt       float64 `json:"created_at,omitempty"`
	Destination     string  `json:"destination,omitempty"`
	Msatoshi        int     `json:"msatoshi,omitempty"`
	MsatoshiSent    int     `json:"msatoshi_sent,omitempty"`
	Parts           int     `json:"parts,omitempty"`
	PaymentHash     string  `json:"payment_hash,omitempty"`
	PaymentPreimage string  `json:"payment_preimage,omitempty"`
	Status          string  `json:"status,omitempty"`

	ResponseError
	Error *ResponseError `json:"error,omitempty"`
}
