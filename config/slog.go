package config

import (
	"log/slog"
	"os"
)

func NewSlogger(debug bool) *slog.Logger {
	level := slog.LevelWarn
	if debug {
		level = slog.LevelDebug
	}

	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level, AddSource: true}))
}
