package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
)

type customLogHandler struct {
	slog.Handler
}

func (h *customLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level == slog.LevelError {
		// capture stack trace (skip first few frames)
		stack := debug.Stack()
		r.AddAttrs(slog.String("stack", string(stack)))
	}
	return h.Handler.Handle(ctx, r)
}

func NewLogger() *slog.Logger {
	base := slog.NewJSONHandler(os.Stdout, nil)
	return slog.New(&customLogHandler{base})
}
