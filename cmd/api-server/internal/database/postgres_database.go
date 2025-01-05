package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type PostgresDatabaseConnector struct{}

type PostgresDatabase struct {
	db *sql.DB
}

var ErrConnectionFailed = errors.New("failed to create database connection")
var ErrDatabaseNotReachable = errors.New("database not reachable")

func (f *PostgresDatabaseConnector) Connect(connectionString string) (Database, error) {
	if db, err := sql.Open("postgres", connectionString); err != nil {
		return nil, ErrConnectionFailed
	} else {
		postgresDb := &PostgresDatabase{db: db}
		if err := postgresDb.Ping(); err != nil {
			return postgresDb, ErrDatabaseNotReachable
		} else {
			return postgresDb, nil
		}
	}
}

func (d *PostgresDatabase) Ping() error {
	return d.db.Ping()
}

func (d *PostgresDatabase) Close() error {
	return d.db.Close()
}
