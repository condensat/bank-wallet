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

func GetInfo(ctx context.Context) (common.GetInfoResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.GetInfoResponse{}, ErrInternalError
	}

	cmd.verb = http.MethodGet
	cmd.command = CommandGetInfo

	data, err := callCommand(cmd)
	if err != nil {
		return common.GetInfoResponse{}, err
	}

	var resp common.GetInfoResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.GetInfoResponse{}, err
	}
	return resp, nil
}
