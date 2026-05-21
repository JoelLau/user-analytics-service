package server

import "context"

var _ StrictServerInterface = &Server{}

func NewServer() *Server {
	return &Server{}
}

type Server struct{}

// (GET /api/livez)
func (s *Server) GetApiLivez(ctx context.Context, request GetApiLivezRequestObject) (GetApiLivezResponseObject, error) {
	return GetApiLivez200JSONResponse{}, nil
}

// (GET /api/readyz)
func (s *Server) GetApiReadyz(ctx context.Context, request GetApiReadyzRequestObject) (GetApiReadyzResponseObject, error) {
	return GetApiReadyz200JSONResponse{}, nil
}

// (GET /api/v1/analytics/users/daily/{day})
func (s *Server) GetApiV1AnalyticsUsersDailyDay(ctx context.Context, request GetApiV1AnalyticsUsersDailyDayRequestObject) (GetApiV1AnalyticsUsersDailyDayResponseObject, error) {
	return GetApiV1AnalyticsUsersDailyDay200JSONResponse{}, nil
}

// (GET /api/v1/analytics/users/monthly/{month})
func (s *Server) GetApiV1AnalyticsUsersMonthlyMonth(ctx context.Context, request GetApiV1AnalyticsUsersMonthlyMonthRequestObject) (GetApiV1AnalyticsUsersMonthlyMonthResponseObject, error) {
	return GetApiV1AnalyticsUsersMonthlyMonth200JSONResponse{}, nil
}

// (POST /api/v1/logins)
func (s *Server) PostApiV1Logins(ctx context.Context, request PostApiV1LoginsRequestObject) (PostApiV1LoginsResponseObject, error) {
	return PostApiV1Logins201Response{}, nil
}
