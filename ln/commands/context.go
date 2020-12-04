// Copyright 2020 Condensat Tech. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"context"
)

const (
	CommandKey = "Key.CommandKey"
)

// WithCommand returns a context with the messaging set
func WithCommand(ctx context.Context, cmd *Command) context.Context {
	return context.WithValue(ctx, CommandKey, cmd)
}

func FromContext(ctx context.Context) *Command {
	if ctxCommand, ok := ctx.Value(CommandKey).(*Command); ok {
		return ctxCommand
	}
	return nil
}
