// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package common

import (
	"testing"
)

func TestIssuanceInfo_IsValid_Mode(t *testing.T) {
	t.Parallel()

	type fields struct {
		Mode               AssetIssuanceMode
		AssetPublicAddress string
		AssetIssuedAmount  float64
		TokenPublicAddress string
		TokenIssuedAmount  float64
		ContractHash       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{}, false},
		{"invalid", fields{Mode: "foobar"}, false},

		{"issueAsset", fields{Mode: AssetIssuanceModeWithAsset, AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1}, true},
		{"issueAssetWithToken", fields{Mode: AssetIssuanceModeWithToken, AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1}, true},
		{"issueAssetWithContract", fields{Mode: AssetIssuanceModeWithContract, AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, ContractHash: "contract"}, true},
		{"issueAssetWithTokenWithContract", fields{Mode: AssetIssuanceModeWithTokenWithContract, AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1, ContractHash: "contract"}, true},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			p := &IssuanceRequest{
				Mode:               tt.fields.Mode,
				AssetPublicAddress: tt.fields.AssetPublicAddress,
				AssetIssuedAmount:  tt.fields.AssetIssuedAmount,
				TokenPublicAddress: tt.fields.TokenPublicAddress,
				TokenIssuedAmount:  tt.fields.TokenIssuedAmount,
				ContractHash:       tt.fields.ContractHash,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IssuanceInfo.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssuanceInfo_IsValid_WithAsset(t *testing.T) {
	t.Parallel()

	type fields struct {
		AssetPublicAddress string
		AssetIssuedAmount  float64
		TokenPublicAddress string
		TokenIssuedAmount  float64
		ContractHash       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{}, false},
		{"invalidAddress", fields{AssetIssuedAmount: 42.1}, false},
		{"invalidAmount", fields{AssetPublicAddress: "foobar"}, false},

		{"invalidContract", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, ContractHash: "invalid"}, false},

		{"valid", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1}, true},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			p := &IssuanceRequest{
				Mode:               AssetIssuanceModeWithAsset,
				AssetPublicAddress: tt.fields.AssetPublicAddress,
				AssetIssuedAmount:  tt.fields.AssetIssuedAmount,
				ContractHash:       tt.fields.ContractHash,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IssuanceInfo.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssuanceInfo_IsValid_WithToken(t *testing.T) {
	t.Parallel()
	type fields struct {
		Mode               AssetIssuanceMode
		AssetPublicAddress string
		AssetIssuedAmount  float64
		TokenPublicAddress string
		TokenIssuedAmount  float64
		ContractHash       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{}, false},
		{"invalidAssetAddress", fields{AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1}, false},
		{"invalidAssetAmount", fields{AssetPublicAddress: "foobar", TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1}, false},
		{"invalidTokenAddress", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenIssuedAmount: 42.1}, false},
		{"invalidTokenAmount", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar"}, false},

		{"invalidContractHash", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1, ContractHash: "contract"}, false},

		{"valid", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IssuanceRequest{
				Mode:               AssetIssuanceModeWithToken,
				AssetPublicAddress: tt.fields.AssetPublicAddress,
				AssetIssuedAmount:  tt.fields.AssetIssuedAmount,
				TokenPublicAddress: tt.fields.TokenPublicAddress,
				TokenIssuedAmount:  tt.fields.TokenIssuedAmount,
				ContractHash:       tt.fields.ContractHash,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IssuanceInfo.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssuanceInfo_IsValid_WithContract(t *testing.T) {
	type fields struct {
		Mode               AssetIssuanceMode
		AssetPublicAddress string
		AssetIssuedAmount  float64
		TokenPublicAddress string
		TokenIssuedAmount  float64
		ContractHash       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{}, false},
		{"invalidAssetAddress", fields{AssetIssuedAmount: 42.1, ContractHash: "contract"}, false},
		{"invalidAssetAmount", fields{AssetPublicAddress: "foobar", ContractHash: "contract"}, false},
		{"invalidContractHash", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1}, false},

		{"invalidTokenAddress", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenIssuedAmount: 42.1, ContractHash: "contract"}, false},
		{"invalidTokenAmount", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", ContractHash: "contract"}, false},

		{"valid", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, ContractHash: "contract"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IssuanceRequest{
				Mode:               AssetIssuanceModeWithContract,
				AssetPublicAddress: tt.fields.AssetPublicAddress,
				AssetIssuedAmount:  tt.fields.AssetIssuedAmount,
				TokenPublicAddress: tt.fields.TokenPublicAddress,
				TokenIssuedAmount:  tt.fields.TokenIssuedAmount,
				ContractHash:       tt.fields.ContractHash,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IssuanceInfo.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssuanceInfo_IsValid_WithTokenWithContract(t *testing.T) {
	type fields struct {
		Mode               AssetIssuanceMode
		AssetPublicAddress string
		AssetIssuedAmount  float64
		TokenPublicAddress string
		TokenIssuedAmount  float64
		ContractHash       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{}, false},
		{"invalidAssetAddress", fields{AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1, ContractHash: "contract"}, false},
		{"invalidAssetAmount", fields{AssetPublicAddress: "foobar", TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1, ContractHash: "contract"}, false},
		{"invalidTokenAddress", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenIssuedAmount: 42.1, ContractHash: "contract"}, false},
		{"invalidTokenAmount", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", ContractHash: "contract"}, false},
		{"invalidContractHash", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1}, false},

		{"valid", fields{AssetPublicAddress: "foobar", AssetIssuedAmount: 42.1, TokenPublicAddress: "foobar", TokenIssuedAmount: 42.1, ContractHash: "contract"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IssuanceRequest{
				Mode:               AssetIssuanceModeWithTokenWithContract,
				AssetPublicAddress: tt.fields.AssetPublicAddress,
				AssetIssuedAmount:  tt.fields.AssetIssuedAmount,
				TokenPublicAddress: tt.fields.TokenPublicAddress,
				TokenIssuedAmount:  tt.fields.TokenIssuedAmount,
				ContractHash:       tt.fields.ContractHash,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IssuanceInfo.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
