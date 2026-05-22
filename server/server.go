package server

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ StrictServerInterface = &Server{}

func NewServer(pool *pgxpool.Pool) *Server {
	return &Server{
		pool: pool,
	}
}

type Server struct {
	pool *pgxpool.Pool
}

var ok = "ok"

// (GET /api/livez)
func (s *Server) Livez(ctx context.Context, request LivezRequestObject) (LivezResponseObject, error) {
	return Livez200JSONResponse{Data: &ok}, nil
}

// (GET /api/readyz)
func (s *Server) Readyz(ctx context.Context, request ReadyzRequestObject) (ReadyzResponseObject, error) {
	if err := s.pool.Ping(ctx); err != nil {
		return Readyz500JSONResponse{
			Type:   "https://github.com/JoelLau/user-analytics-service/errors/database-unavailable", // NOTE: link to documentation (doesn't exist)
			Title:  "Service Unavailable",
			Status: http.StatusInternalServerError,
		}, nil
	}

	return Readyz200JSONResponse{Data: &ok}, nil
}

// (GET /api/v1/analytics/users/daily/{day})
func (s *Server) GetDailyUniqueUsers(ctx context.Context, request GetDailyUniqueUsersRequestObject) (GetDailyUniqueUsersResponseObject, error) {
	return GetDailyUniqueUsers200JSONResponse{}, nil

}

// (GET /api/v1/analytics/users/monthly/{month})
func (s *Server) GetMonthlyUniqueUsers(ctx context.Context, request GetMonthlyUniqueUsersRequestObject) (GetMonthlyUniqueUsersResponseObject, error) {
	return GetMonthlyUniqueUsers200JSONResponse{}, nil

}
