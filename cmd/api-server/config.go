package main

import (
	"fmt"

	"github.com/milesjpool/kahd/internal/env"
)

var DATABASE_CONNECTION_STRING_KEY = "DATABASE_CONNECTION_STRING"
var PORT_KEY = "PORT"
var DEFAULT_PORT = "8080"

type Config struct {
	Port         string
	DBConnection string
}

func loadConfig() (Config, error) {
	dbConnection, err := env.Get(DATABASE_CONNECTION_STRING_KEY)
	if err != nil {
		return Config{}, fmt.Errorf("failed to get database URL: %w", err)
	}

	return Config{
		Port:         env.GetOrDefault(PORT_KEY, DEFAULT_PORT),
		DBConnection: dbConnection,
	}, nil
}
