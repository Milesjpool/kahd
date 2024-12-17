package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedPort = "123123123"
var expectedDBConnection = "my-db-connection-string"

func setupTest(t *testing.T) {
	t.Helper()
	os.Setenv(DATABASE_CONNECTION_STRING_KEY, expectedDBConnection)
	os.Setenv(PORT_KEY, expectedPort)
}

func Test_loadConfig(t *testing.T) {

	t.Run("it retrieves config from env", func(t *testing.T) {
		setupTest(t)
		config, err := loadConfig()

		assert.NoError(t, err)
		assert.Equal(t, config.DBConnection, expectedDBConnection)
		assert.Equal(t, config.Port, expectedPort)
	})

	t.Run("it errors if the DB connection string isn't set", func(t *testing.T) {
		setupTest(t)
		os.Unsetenv(DATABASE_CONNECTION_STRING_KEY)
		_, err := loadConfig()

		assert.ErrorContains(t, err, DATABASE_CONNECTION_STRING_KEY)
	})

	t.Run("it uses a default if the port isn't set", func(t *testing.T) {
		setupTest(t)
		os.Unsetenv(PORT_KEY)
		config, err := loadConfig()

		assert.NoError(t, err)
		assert.Equal(t, config.Port, DEFAULT_PORT)
	})
}
