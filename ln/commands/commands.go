// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	CommandGetInfo      = "/v1/getinfo"
	CommandKeySend      = "/v1/pay/keysend"
	CommandPay          = "/v1/pay"
	CommandDecodePay    = "/v1/pay/decodePay"
	CommandListInvoices = "/v1/invoice/listInvoices"
)

var (
	ErrInternalError  = errors.New("Internal Error")
	ErrInvalidPubKey  = errors.New("Invalid PubKey")
	ErrInvalidAmount  = errors.New("Invalid Amount")
	ErrInvalidInvoice = errors.New("Invalid Invoice")
)

type Command struct {
	client   *http.Client
	endpoint string
	macaroon string
	verb     string
	command  string
	isJson   bool
	body     io.Reader
}

func NewCommand(client *http.Client, endpoint string) *Command {
	return &Command{
		client:   client,
		endpoint: endpoint,
	}
}

func (p *Command) WithMacaroon(macaroon string) *Command {
	resp := *p
	resp.macaroon = macaroon
	return &resp
}

func (p *Command) WithBody(body io.Reader) *Command {
	if body == nil {
		return p
	}

	resp := *p
	resp.body = body
	return &resp
}

func (p *Command) WithArgs(args interface{}) *Command {
	data, err := json.Marshal(args)
	if err != nil {
		return p
	}

	resp := *p
	resp.isJson = true

	return resp.WithBody(
		bytes.NewReader(data),
	)
}
