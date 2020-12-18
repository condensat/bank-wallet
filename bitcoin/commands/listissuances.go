// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import "context"

func ListIssuances(ctx context.Context, rpcClient RpcClient, asset AssetID) ([]ListIssuancesInfo, error) {
	var result []ListIssuancesInfo

	err := callCommand(rpcClient, CmdListIssuances, &result, asset)
	if err != nil {
		return nil, err
	}
	return result, nil
}
