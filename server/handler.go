package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(slogger *slog.Logger) http.Handler {
	serverImpl := NewServer()
	strictHandler := NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()

	r.Use(slogchi.New(slogger))
	r.Use(middleware.Recoverer)

	HandlerFromMux(strictHandler, r)

	return r
}
