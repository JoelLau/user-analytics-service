package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(slogger *slog.Logger, pool *pgxpool.Pool) http.Handler {
	serverImpl := NewServer(pool)
	strictHandler := NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()

	r.Use(slogchi.New(slogger))
	r.Use(middleware.Recoverer)

	HandlerFromMux(strictHandler, r)

	return r
}
