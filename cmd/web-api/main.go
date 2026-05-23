package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-analytics/config"
	"user-analytics/queries"
	"user-analytics/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

const shutdownTimeout = 10 * time.Second

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Init(ctx)
	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting..")

	pool, err := pgxpool.New(ctx, cfg.DSN())
	if err != nil {
		logr.ErrorContext(ctx, "failed to create db pool", slog.Any("error", err))
		return
	}
	defer pool.Close()

	srv := &http.Server{
		Handler: server.NewHandler(logr, queries.New(pool), NewClock()),
		Addr:    cfg.Address,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logr.ErrorContext(ctx, "unexpected server shutdown", slog.Any("error", err))
		}
	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logr.ErrorContext(shutdownCtx, "server shutdown error", slog.Any("error", err))
	}

	logr.InfoContext(shutdownCtx, "exiting..")
}

type Clock struct{}

func NewClock() *Clock {
	return &Clock{}
}

func (c *Clock) Now() time.Time {
	return time.Now()
}
