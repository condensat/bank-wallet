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

func KeySend(ctx context.Context, pubKey string, amount int, label string) (common.KeySendResponse, error) {
	cmd := FromContext(ctx)
	if cmd == nil {
		return common.KeySendResponse{}, ErrInternalError
	}
	if len(pubKey) == 0 {
		return common.KeySendResponse{}, ErrInvalidPubKey
	}
	if amount <= 0 {
		return common.KeySendResponse{}, ErrInvalidAmount
	}

	cmd.verb = http.MethodPost
	cmd.command = CommandKeySend
	cmd = cmd.WithArgs(&common.KeySendArgs{
		PubKey: pubKey,
		Amount: amount,
		Label:  label,
	})

	data, err := callCommand(cmd)
	if err != nil {
		return common.KeySendResponse{}, err
	}

	var resp common.KeySendResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return common.KeySendResponse{}, err
	}

	return resp, nil
}
