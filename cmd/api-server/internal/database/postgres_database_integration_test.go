//go:build integration

package database

import (
	"testing"

	"github.com/milesjpool/kahd/internal/env"
	"github.com/stretchr/testify/assert"
)

func getConnectionString(t *testing.T) string {
	t.Helper()

	connStr, err := env.Get("TEST_DB_CONNECTION_STRING")
	if err != nil {
		t.Fatalf("failed to get TEST_DB_CONNECTION_STRING: %v", err)
	}
	return connStr
}

func TestPostgresDatabaseConnector_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	connStr := getConnectionString(t)
	connector := &PostgresDatabaseConnector{}

	t.Run("successfully connects to database", func(t *testing.T) {
		db, err := connector.Connect(connStr)
		assert.NoError(t, err)
		defer db.Close()

		err = db.Ping()
		assert.NoError(t, err)
	})

	t.Run("ping errors if connection is already closed", func(t *testing.T) {
		db, err := connector.Connect(connStr)
		assert.NoError(t, err)
		db.Close()

		err = db.Ping()
		assert.Error(t, err)
	})

	t.Run("close does not error if connection is already closed", func(t *testing.T) {
		db, err := connector.Connect(connStr)
		assert.NoError(t, err)
		db.Close()

		err = db.Close()
		assert.NoError(t, err)
	})

	t.Run("returns error on invalid connection string", func(t *testing.T) {
		_, err := connector.Connect("invalid connection string")
		assert.Error(t, err)
	})
}
