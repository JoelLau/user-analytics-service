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

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: true}))
}
