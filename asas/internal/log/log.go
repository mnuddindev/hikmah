package log

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(h)
}