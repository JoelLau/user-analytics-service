package server

import (
	"context"
	"net/http"
	"time"
	"user-analytics/queries"
)

var _ StrictServerInterface = &Server{}

func NewServer(q *queries.Queries, clk Nower) *Server {
	return &Server{queries: q, clock: clk}
}

type Server struct {
	queries *queries.Queries
	clock   Nower
}

// (GET /api/livez)
func (s *Server) Livez(ctx context.Context, request LivezRequestObject) (LivezResponseObject, error) {
	return Livez200JSONResponse{Data: new("ok")}, nil
}

// (GET /api/readyz)
func (s *Server) Readyz(ctx context.Context, request ReadyzRequestObject) (ReadyzResponseObject, error) {
	if err := s.queries.Ping(ctx); err != nil {
		return Readyz500JSONResponse{
			Type:   "https://github.com/JoelLau/user-analytics-service/errors/database-unavailable",
			Title:  "Service Unavailable",
			Status: http.StatusInternalServerError,
		}, nil
	}

	return Readyz200JSONResponse{Data: new("ok")}, nil
}

// (GET /api/v1/analytics/users/daily/{day})
func (s *Server) GetDailyUniqueUsers(ctx context.Context, request GetDailyUniqueUsersRequestObject) (GetDailyUniqueUsersResponseObject, error) {
	// e.g. 2026-05-22 00:00:00 UTC
	day := request.Day.Time.UTC().Truncate(24 * time.Hour)
	now := s.clock.Now().UTC()

	if day.After(now) {
		return GetDailyUniqueUsers400JSONResponse{
			Type:   "https://github.com/JoelLau/user-analytics-service/errors/invalid-param",
			Title:  "Bad Request",
			Status: http.StatusBadRequest,
			Detail: new("day must not be in the future"),
		}, nil
	}

	count, err := s.queries.GetDailyUniqueUsers(ctx, day)
	if err != nil {
		return nil, err
	}

	return GetDailyUniqueUsers200JSONResponse{Data: new(int(count))}, nil
}

// (GET /api/v1/analytics/users/monthly/{month})
func (s *Server) GetMonthlyUniqueUsers(ctx context.Context, request GetMonthlyUniqueUsersRequestObject) (GetMonthlyUniqueUsersResponseObject, error) {
	return GetMonthlyUniqueUsers200JSONResponse{}, nil
}
