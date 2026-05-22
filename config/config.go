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
	PostgresUser:     "postgres",
	PostgresPassword: "password",
	PostgresDb:       "postgres",
	PostgresPort:     5432,
}

func Init(ctx context.Context) Config {
	cfg := DefaultConfig

	if err := env.Feed(&cfg); err != nil {
		panic(fmt.Errorf("failed to marshal envvars into struct:%w", err))
	}

	return cfg
}

type Config struct {
	Address          string `env:"ADDRESS"`           // what address to serve the HTTP server           e.g. ":8080"
	DebugMode        bool   `env:"DEBUG"`             // controls log verbosity                          e.g. false
	PostgresUser     string `env:"POSTGRES_USER"`     // (used for db connection) postgres user          e.g. "john.doe"
	PostgresPassword string `env:"POSTGRES_PASSWORD"` // (used for db connection) postgres password      e.g. "p@ssw0rd!123"
	PostgresDb       string `env:"POSTGRES_DB"`       // (used for db connection) postgres database name e.g. "user_analytics"
	PostgresPort     int    `env:"POSTGRES_PORT"`     // (used for db connection) postgres port          e.g. "5432"
}

func (cfg *Config) Logger() *slog.Logger {
	return NewSlogger(cfg.DebugMode)
}
