package server

import (
	"context"
	"net/http"
	"user-analytics/queries"
)

var _ StrictServerInterface = &Server{}

func NewServer(q *queries.Queries) *Server {
	return &Server{queries: q}
}

type Server struct {
	queries *queries.Queries
}

// (GET /api/livez)
func (s *Server) Livez(ctx context.Context, request LivezRequestObject) (LivezResponseObject, error) {
	return Livez200JSONResponse{Data: ptr("ok")}, nil
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

	return Readyz200JSONResponse{Data: ptr("ok")}, nil
}

// (GET /api/v1/analytics/users/daily/{day})
func (s *Server) GetDailyUniqueUsers(ctx context.Context, request GetDailyUniqueUsersRequestObject) (GetDailyUniqueUsersResponseObject, error) {
	day := request.Day.Time

	count, err := s.queries.GetDailyUniqueUsers(ctx, day)
	if err != nil {
		return nil, err
	}

	return GetDailyUniqueUsers200JSONResponse{Data: ptr(int(count))}, nil
}

// (GET /api/v1/analytics/users/monthly/{month})
func (s *Server) GetMonthlyUniqueUsers(ctx context.Context, request GetMonthlyUniqueUsersRequestObject) (GetMonthlyUniqueUsersResponseObject, error) {
	return GetMonthlyUniqueUsers200JSONResponse{}, nil
}

func ptr[T any](v T) *T { return new(v) }
