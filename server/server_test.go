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
	srv := httptest.NewServer(server.NewHandler(discardLogger, nil))
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
	srv, _ := newTestServer(t)

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
	srv := httptest.NewServer(server.NewHandler(discardLogger, q))
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
	srv, pool := newTestServer(t)

	insertUserLogins(t, pool, []queries.InsertUserLoginParams{
		{UserID: 1, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 8, 0, 0, 0, time.UTC)},
		{UserID: 2, LoggedInAt: time.Date(2026, 3, 15, 20, 0, 0, 0, time.UTC)}, // same user same day — must not double-count
		{UserID: 3, LoggedInAt: time.Date(2026, 3, 16, 8, 0, 0, 0, time.UTC)},  // different day — must be excluded
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
