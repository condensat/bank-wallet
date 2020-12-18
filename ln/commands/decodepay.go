// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/condensat/bank-wallet/common"
)

func DecodePay(ctx context.Context, invoice string) (common.DecodePayResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.DecodePayResponse{}, ErrInternalError
	}
	if len(invoice) == 0 {
		return common.DecodePayResponse{}, ErrInvalidInvoice
	}

	cmd.verb = http.MethodGet
	cmd.command = fmt.Sprintf("%s/%s", CommandDecodePay, invoice)

	data, err := callCommand(cmd)
	if err != nil {
		return common.DecodePayResponse{}, err
	}

	var resp common.DecodePayResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.DecodePayResponse{}, err
	}

	return resp, nil
}
