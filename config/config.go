package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/golobby/env/v2"
)

var DefaultConfig = Config{
	Address:          ":8080",
	DebugMode:        false,
	PostgresHost:     "localhost",
	PostgresUser:     "postgres",
	PostgresPassword: "password",
	PostgresDb:       "postgres",
	PostgresPort:     5432,
}

// WARN: replaces global default logger
func Init(ctx context.Context) Config {
	cfg := DefaultConfig

	if err := env.Feed(&cfg); err != nil {
		panic(fmt.Errorf("failed to marshal envvars into struct:%w", err))
	}

	slog.InfoContext(ctx, "env-ed", slog.Any("cfg", cfg))

	return cfg
}

type Config struct {
	Address          string `env:"ADDRESS"`           // what address to serve the HTTP server           e.g. ":8080"
	DebugMode        bool   `env:"DEBUG"`             // controls log verbosity                          e.g. false
	PostgresHost     string `env:"POSTGRES_HOST"`     // (used for db connection) postgres host          e.g. "localhost"
	PostgresUser     string `env:"POSTGRES_USER"`     // (used for db connection) postgres user          e.g. "john.doe"
	PostgresPassword string `env:"POSTGRES_PASSWORD"` // (used for db connection) postgres password      e.g. "p@ssw0rd!123"
	PostgresDb       string `env:"POSTGRES_DB"`       // (used for db connection) postgres database name e.g. "user_analytics"
	PostgresPort     int    `env:"POSTGRES_PORT"`     // (used for db connection) postgres port          e.g. "5432"
}

func (cfg *Config) Logger() *slog.Logger {
	return NewSlogger(cfg.DebugMode)
}

func (cfg *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb, cfg.PostgresPort)
}
