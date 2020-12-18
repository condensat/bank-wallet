// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ln

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync"

	"github.com/condensat/bank-wallet/common"
	"github.com/condensat/bank-wallet/ln/commands"
)

var (
	ErrKeySendError = errors.New("KeySend Error")
)

type LightningClient struct {
	sync.Mutex

	client   *http.Client
	endpoint string
	macaroon string
}

func NewWithTorEndpoint(ctx context.Context, torProxy, endpoint, macaroon string) *LightningClient {
	proxyURL, err := url.Parse(torProxy)
	if err != nil {
		return nil
	}

	return &LightningClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		},
		endpoint: endpoint,
		macaroon: macaroon,
	}
}

func defaultCommand(ln *LightningClient) *commands.Command {
	return commands.
		NewCommand(ln.client, ln.endpoint).
		WithMacaroon(ln.macaroon)
}

func (p *LightningClient) GetInfo(ctx context.Context) (common.GetInfoResponse, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	ctx = commands.WithCommand(ctx, defaultCommand(p))

	return commands.GetInfo(ctx)
}

func (p *LightningClient) KeySend(ctx context.Context, pubKey string, amount int, label string) (common.KeySendResponse, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	ctx = commands.WithCommand(ctx, defaultCommand(p))

	resp, err := commands.KeySend(ctx, pubKey, amount, label)
	if err != nil {
		return common.KeySendResponse{}, err
	}

	if resp.Code != 0 {
		return common.KeySendResponse{
			ResponseError: resp.ResponseError,
		}, ErrKeySendError
	}

	return resp, nil
}

func (p *LightningClient) Pay(ctx context.Context, invoice string) (common.PayResponse, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	ctx = commands.WithCommand(ctx, defaultCommand(p))

	resp, err := commands.Pay(ctx, invoice)
	if err != nil {
		return common.PayResponse{}, err
	}

	if resp.Code != 0 {
		return common.PayResponse{
			ResponseError: resp.ResponseError,
		}, ErrKeySendError
	}

	return resp, nil
}

func (p *LightningClient) DecodePay(ctx context.Context, invoice string) (common.DecodePayResponse, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	ctx = commands.WithCommand(ctx, defaultCommand(p))

	resp, err := commands.DecodePay(ctx, invoice)
	if err != nil {
		return common.DecodePayResponse{}, err
	}

	if resp.Code != 0 {
		return common.DecodePayResponse{
			ResponseError: resp.ResponseError,
		}, ErrKeySendError
	}

	return resp, nil
}
