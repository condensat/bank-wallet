// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/condensat/bank-wallet/common"
)

func ListInvoices(ctx context.Context, label string) (common.ListInvoicesResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.ListInvoicesResponse{}, ErrInternalError
	}

	cmd.verb = http.MethodGet
	cmd.command = CommandListInvoices
	if len(label) != 0 {
		cmd.command = fmt.Sprintf("%s/?label=%s", CommandListInvoices, url.QueryEscape(label))
	}

	data, err := callCommand(cmd)
	if err != nil {
		return common.ListInvoicesResponse{}, err
	}

	var resp common.ListInvoicesResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.ListInvoicesResponse{}, err
	}

	return resp, nil
}
