// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/condensat/bank-wallet/common"
)

func GenInvoice(ctx context.Context, amount int, label, description string, expiry int, private bool) (common.GenInvoiceResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.GenInvoiceResponse{}, ErrInternalError
	}

	cmd.verb = http.MethodPost
	cmd.command = CommandGenInvoice
	cmd = cmd.WithArgs(&common.GenInvoiceArgs{
		Amount:      amount,
		Label:       label,
		Description: description,
		Expiry:      expiry,
		Private:     private,
	})

	data, err := callCommand(cmd)
	if err != nil {
		return common.GenInvoiceResponse{}, err
	}

	var resp common.GenInvoiceResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.GenInvoiceResponse{}, err
	}

	return resp, nil
}
