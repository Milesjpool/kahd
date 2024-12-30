package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

var unreachableDatabase = "postgres://abc:def@unknown:123/db"

func TestPostgresDatabaseConnector(t *testing.T) {
	t.Run("it returns an error if the database is not reachable", func(t *testing.T) {
		connector := &PostgresDatabaseConnector{}
		db, err := connector.Connect(unreachableDatabase)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}

func TestPostgresDatabase(t *testing.T) {
	t.Run("ping returns an error if the database is not reachable", func(t *testing.T) {
		db, err := sql.Open("postgres", unreachableDatabase)
		assert.NoError(t, err)

		postgresDatabase := &PostgresDatabase{db: db}

		err = postgresDatabase.Ping()
		assert.Error(t, err)
	})
}
