package server

import (
	"context"
)

var _ StrictServerInterface = &Server{}

func NewServer() *Server {
	return &Server{}
}

type Server struct{}

// (GET /api/livez)
func (s *Server) Livez(ctx context.Context, request LivezRequestObject) (LivezResponseObject, error) {
	return Livez200JSONResponse{}, nil

}

// (GET /api/readyz)
func (s *Server) Readyz(ctx context.Context, request ReadyzRequestObject) (ReadyzResponseObject, error) {
	return Readyz200JSONResponse{}, nil
}

// (GET /api/v1/analytics/users/daily/{day})
func (s *Server) GetDailyUniqueUsers(ctx context.Context, request GetDailyUniqueUsersRequestObject) (GetDailyUniqueUsersResponseObject, error) {
	return GetDailyUniqueUsers200JSONResponse{}, nil

}

// (GET /api/v1/analytics/users/monthly/{month})
func (s *Server) GetMonthlyUniqueUsers(ctx context.Context, request GetMonthlyUniqueUsersRequestObject) (GetMonthlyUniqueUsersResponseObject, error) {
	return GetMonthlyUniqueUsers200JSONResponse{}, nil

}

// (POST /api/v1/logins)
func (s *Server) RecordLogin(ctx context.Context, request RecordLoginRequestObject) (RecordLoginResponseObject, error) {
	return RecordLogin201Response{}, nil
}
