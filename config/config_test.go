package config_test

import (
	"testing"
	"user-analytics/config"

	"github.com/stretchr/testify/assert"
)

func TestConfig_DSN(t *testing.T) {
	cfg := config.Config{
		PostgresHost:     "localhost",
		PostgresUser:     "john.doe",
		PostgresPassword: "p@ssw0rd!123",
		PostgresDb:       "user-analytics",
		PostgresPort:     5432,
	}

	assert.Equal(t,
		"host=localhost user=john.doe password=p@ssw0rd!123 dbname=user-analytics port=5432 sslmode=disable",
		cfg.DSN(),
	)
}
