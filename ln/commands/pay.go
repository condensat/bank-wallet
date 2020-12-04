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

func Pay(ctx context.Context, invoice string) (common.PayResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.PayResponse{}, ErrInternalError
	}
	if len(invoice) == 0 {
		return common.PayResponse{}, ErrInvalidInvoice
	}

	cmd.verb = http.MethodPost
	cmd.command = CommandPay
	cmd = cmd.WithArgs(&common.PayArgs{
		Invoice: invoice,
	})

	data, err := callCommand(cmd)
	if err != nil {
		return common.PayResponse{}, err
	}

	var resp common.PayResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.PayResponse{}, err
	}

	return resp, nil
}
