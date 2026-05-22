package server_test

import (
	"context"
	"log/slog"
	"net/http/httptest"
	"testing"
	"user-analytics/server"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var discardLogger = slog.New(slog.DiscardHandler)

func newPgContainer(t testing.TB, ctx context.Context) *postgres.PostgresContainer {
	pgCont, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = pgCont.Terminate(context.Background()) })

	return pgCont
}

func newPool(t testing.TB, ctx context.Context, connStr string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connStr)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	return pool
}

func newTestServer(t testing.TB) *httptest.Server {
	pgCont := newPgContainer(t, t.Context())
	connStr, err := pgCont.ConnectionString(t.Context(), "sslmode=disable")
	require.NoError(t, err)

	pool := newPool(t, t.Context(), connStr)
	srv := httptest.NewServer(server.NewHandler(discardLogger, pool))
	t.Cleanup(srv.Close)

	return srv
}
