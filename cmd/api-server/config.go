package main

import (
	"fmt"

	"github.com/milesjpool/kahd/internal/env"
)

var DATABASE_CONNECTION_STRING_KEY = "DATABASE_CONNECTION_STRING"
var PORT_KEY = "PORT"
var DEFAULT_PORT = "8080"

type EnvConfigLoader struct{}

type APIConfig struct {
	Port         string
	DBConnection string
}

func (c *EnvConfigLoader) Load() (APIConfig, error) {
	dbConnection, err := env.Get(DATABASE_CONNECTION_STRING_KEY)
	if err != nil {
		return APIConfig{}, fmt.Errorf("failed to get database URL: %w", err)
	}

	return APIConfig{
		Port:         env.GetOrDefault(PORT_KEY, DEFAULT_PORT),
		DBConnection: dbConnection,
	}, nil
}
