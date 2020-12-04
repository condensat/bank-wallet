// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func callCommand(cmd *Command) ([]byte, error) {
	if cmd == nil {
		return nil, ErrInternalError
	}
	if cmd.client == nil {
		return nil, ErrInternalError
	}
	if len(cmd.endpoint) == 0 {
		return nil, ErrInternalError
	}
	if len(cmd.command) == 0 {
		return nil, ErrInternalError
	}

	url := fmt.Sprintf("%s%s", cmd.endpoint, cmd.command)

	// create request
	req, err := http.NewRequest(cmd.verb, url, cmd.body)
	if err != nil {
		return nil, err
	}

	// set macaroon header
	if len(cmd.macaroon) > 0 {
		req.Header.Set("encodingtype", "hex")
		req.Header.Set("macaroon", cmd.macaroon)
	}

	if cmd.isJson {
		if cmd.body == nil {
			panic("Nil JSON Body")
		}
		req.Header.Set("Content-type", "application/json")
	}

	// request
	resp, err := cmd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
