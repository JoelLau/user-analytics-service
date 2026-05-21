package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"user-analytics/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()
	cfg := config.Init(ctx)

	slog.InfoContext(ctx, "starting..")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.Address,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "unexpected server shutdown", slog.Any("error", err))
		return
	}

	slog.InfoContext(ctx, "exiting..")
}
