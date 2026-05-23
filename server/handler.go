package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"user-analytics/queries"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(slogger *slog.Logger, q *queries.Queries) http.Handler {
	serverImpl := NewServer(q)
	strictHandler := NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()

	r.Use(slogchi.New(slogger))
	r.Use(middleware.Recoverer)

	HandlerWithOptions(strictHandler, ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: errorHandlerFunc,
	})

	return r
}

var paramExampleByName = map[string]string{
	"day":   "2026-05-21",
	"month": "2026-05",
}

func errorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	var detail string

	switch e := err.(type) {
	case *InvalidParamFormatError:
		detail = fmt.Sprintf("invalid value for '%s': %s (e.g. '%q')", e.ParamName, e.Err.Error(), paramExampleByName[e.ParamName])
	case *RequiredParamError:
		detail = fmt.Sprintf("'%s' is required (e.g. '%q')", e.ParamName, paramExampleByName[e.ParamName])
	default:
		detail = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Type:   "https://github.com/JoelLau/user-analytics-service/errors/invalid-param",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: new(detail),
	})
}
