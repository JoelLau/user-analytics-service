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
	"user-analytics/server"
)

const shutdownTimeout = 10 * time.Second

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Init(ctx)
	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting..")

	srv := &http.Server{
		Handler: server.NewHandler(logr),
		Addr:    cfg.Address,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "unexpected server shutdown", slog.Any("error", err))
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
