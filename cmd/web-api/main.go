package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"user-analytics/config"
	"user-analytics/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

func main() {
	ctx := context.Background()
	cfg := config.Init(ctx)

	slog.InfoContext(ctx, "starting..")

	serverImpl := server.NewServer()
	strictHandler := server.NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()
	r.Use(slogchi.New(slog.Default()))
	r.Use(middleware.Recoverer)

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.Address,
	}

	server.HandlerFromMux(strictHandler, r)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "unexpected server shutdown", slog.Any("error", err))
		return
	}

	slog.InfoContext(ctx, "exiting..")
}
