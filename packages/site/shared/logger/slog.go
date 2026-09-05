package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
	"slices"
	"strings"

	colorjson "github.com/hydronica/color-json"
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
	var handler slog.Handler = colorjson.NewHandler(os.Stdout, nil)
	if val, ok := os.LookupEnv("NO_COLOR"); ok &&
		slices.Contains([]string{"1", "true", ""}, strings.ToLower(val)) {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}

	return slog.New(&customLogHandler{handler})
}
