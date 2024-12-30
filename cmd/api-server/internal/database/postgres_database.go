package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDatabaseConnector struct{}

type PostgresDatabase struct {
	db *sql.DB
}

func (f *PostgresDatabaseConnector) Connect(connectionString string) (Database, error) {
	if db, err := sql.Open("postgres", connectionString); err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	} else if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	} else {
		return &PostgresDatabase{db: db}, nil
	}
}

func (d *PostgresDatabase) Ping() error {
	return d.db.Ping()
}

func (d *PostgresDatabase) Close() error {
	return d.db.Close()
}
