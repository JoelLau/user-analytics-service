package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-analytics/server"

	"github.com/stretchr/testify/require"
)

func TestLivez(t *testing.T) {
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
	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_500InternalServerError(t *testing.T) {
	// Arrange
	pgCont := newPgContainer(t, t.Context())
	connStr, err := pgCont.ConnectionString(t.Context(), "sslmode=disable")
	require.NoError(t, err)

	pool := newPool(t, t.Context(), connStr)
	require.NoError(t, pgCont.Terminate(t.Context()))

	srv := httptest.NewServer(server.NewHandler(discardLogger, pool))
	t.Cleanup(srv.Close)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
