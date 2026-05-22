package server_test

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-analytics/server"

	"github.com/stretchr/testify/require"
)

var discardLogger = slog.New(slog.DiscardHandler)

func TestLivez(t *testing.T) {
	// Arrange
	srv := httptest.NewServer(server.NewHandler(discardLogger))
	defer srv.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/livez", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
