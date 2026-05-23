package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"user-analytics/queries"
	"user-analytics/server"

	"github.com/stretchr/testify/require"
)

func TestLivez(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := httptest.NewServer(server.NewHandler(discardLogger, nil, NewFakeNower(time.Now())))
	defer srv.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/livez", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, _ := newTestServer(t, NewFakeNower(time.Now()))

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_500InternalServerError(t *testing.T) {
	t.Parallel()

	// Arrange
	pgCont := newPgContainer(t, t.Context())
	connStr, err := pgCont.ConnectionString(t.Context(), "sslmode=disable")
	require.NoError(t, err)

	pool := newPool(t, t.Context(), connStr)
	err = pgCont.Terminate(t.Context())
	require.NoError(t, err)

	q := queries.New(pool)
	srv := httptest.NewServer(server.NewHandler(discardLogger, q, NewFakeNower(time.Now())))
	t.Cleanup(srv.Close)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetDailyUniqueUsers_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Now()))

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC)}, // duplicate login
		{UserID: 3, LoggedInAt: time.Date(2026, 3, 16, 8, 0, 0, 0, time.UTC)},  // different day
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/daily/2026-03-15", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var body struct {
		Data int `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	require.NoErrorf(t, err, "failed to decode json response: %w", err)
	require.Equal(t, 2, body.Data)
}

func TestGetDailyUniqueUsers_400BadRequest_InvalidDateFormat(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, _ := newTestServer(t, NewFakeNower(time.Now()))

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/daily/not-a-date", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetDailyUniqueUsers_400BadRequest_FutureDate(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Date(2026, 03, 15, 0, 0, 0, 0, time.UTC))) // hardcoded "now" date

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/daily/2026-03-16", srv.URL)) // hardcoded "future" date
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetMonthlyUniqueUsers_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Now()))

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC)}, // duplicate login
		{UserID: 3, LoggedInAt: time.Date(2026, 4, 16, 8, 0, 0, 0, time.UTC)},  // different month
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/monthly/2026-03", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var body struct {
		Data int `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	require.NoErrorf(t, err, "failed to decode json response: %w", err)
	require.Equal(t, 2, body.Data)
}

func TestGetDailyUniqueUsers_200OK_DuplicateLogins(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Now()))

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC)},
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC)},
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/daily/2026-03-15", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert: same user logging in 3 times counts as 1
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var body struct {
		Data int `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	require.NoErrorf(t, err, "failed to decode json response: %w", err)
	require.Equal(t, 1, body.Data)
}

func TestGetMonthlyUniqueUsers_200OK_DuplicateLogins(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Now()))

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 1, 8, 0, 0, 0, time.UTC)},
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 31, 8, 0, 0, 0, time.UTC)},
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/monthly/2026-03", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert: same user logging in across multiple days in a month counts as 1
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var body struct {
		Data int `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	require.NoErrorf(t, err, "failed to decode json response: %w", err)
	require.Equal(t, 1, body.Data)
}

func TestGetMonthlyUniqueUsers_400BadRequest_InvalidDateFormat(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, _ := newTestServer(t, NewFakeNower(time.Now()))

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/monthly/not-a-date", srv.URL))
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetMonthlyUniqueUsers_400BadRequest_FutureDate(t *testing.T) {
	t.Parallel()

	// Arrange
	srv, pool := newTestServer(t, NewFakeNower(time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC))) // hardcoded "now" date

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC)},
	})

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/analytics/users/monthly/2026-04", srv.URL)) // hardcoded "future" date
	require.NoErrorf(t, err, "error on http GET: %w", err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
