package server_test

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http/httptest"
	"testing"
	"user-analytics/queries"
	"user-analytics/server"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
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
		postgres.BasicWaitStrategies(),
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

func runMigrations(t testing.TB, connStr string) {
	t.Helper()
	db, err := sql.Open("pgx", connStr)
	require.NoError(t, err)
	defer db.Close()

	err = goose.Up(db, "../migrations")
	require.NoError(t, err)
}

func newTestServer(t testing.TB) (*httptest.Server, *pgxpool.Pool) {
	ctx := t.Context()

	pgCont := newPgContainer(t, ctx)
	connStr, err := pgCont.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	runMigrations(t, connStr)

	pool := newPool(t, ctx, connStr)
	q := queries.New(pool)
	srv := httptest.NewServer(server.NewHandler(discardLogger, q))
	t.Cleanup(srv.Close)

	return srv, pool
}

func insertUserLogins(t testing.TB, pool *pgxpool.Pool, logins []queries.InsertUserLoginParams) {
	t.Helper()
	ctx := t.Context()

	tx, err := pool.Begin(ctx)
	require.NoError(t, err)

	q := queries.New(tx)
	for _, p := range logins {
		err = q.InsertUserLogin(ctx, p)
		require.NoError(t, err)
	}

	err = tx.Commit(ctx)
	require.NoError(t, err)

	var count int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM user_logins").Scan(&count)
	require.NoError(t, err)
	require.Equal(t, len(logins), count, "row count after insert")
}
