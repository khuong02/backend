package logger

import (
	"context"
	"golang.org/x/exp/slog"
)

type CustomHandler struct {
	slog.Handler
}

func (c *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	err := c.Handler.Handle(ctx, r)
	if err != nil {
		return err
	}
	return nil
}
