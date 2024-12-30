package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedPort = "123123123"
var expectedDBConnection = "my-db-connection-string"

func Test_EnvConfigLoader(t *testing.T) {

	setupTest := func(t *testing.T) *EnvConfigLoader {
		t.Helper()
		os.Setenv(DATABASE_CONNECTION_STRING_KEY, expectedDBConnection)
		os.Setenv(PORT_KEY, expectedPort)
		return &EnvConfigLoader{}
	}

	t.Run("it retrieves config from env", func(t *testing.T) {
		it := setupTest(t)
		config, err := it.Load()

		assert.NoError(t, err)
		assert.Equal(t, config.DBConnection, expectedDBConnection)
		assert.Equal(t, config.Port, expectedPort)
	})

	t.Run("it errors if the DB connection string isn't set", func(t *testing.T) {
		it := setupTest(t)
		os.Unsetenv(DATABASE_CONNECTION_STRING_KEY)
		_, err := it.Load()

		assert.ErrorContains(t, err, DATABASE_CONNECTION_STRING_KEY)
	})

	t.Run("it uses a default if the port isn't set", func(t *testing.T) {
		it := setupTest(t)
		os.Unsetenv(PORT_KEY)
		config, err := it.Load()

		assert.NoError(t, err)
		assert.Equal(t, config.Port, DEFAULT_PORT)
	})
}
