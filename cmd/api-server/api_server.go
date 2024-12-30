package main

import (
	"fmt"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server"
)

type Loader[T interface{}] interface {
	Load() (T, error)
}
type Connector[T interface{}] interface {
	Connect(connectionString string) (T, error)
}
type ServerFactory[T interface{}] interface {
	NewServer(T) server.Server
}

type APIServer struct {
	ConfigLoader      Loader[APIConfig]
	DatabaseConnector Connector[database.Database]
	ServerFactory     ServerFactory[server.HTTPServerProps]
}

func (s *APIServer) Start() error {
	config, err := s.ConfigLoader.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	db, err := s.DatabaseConnector.Connect(config.DBConnection)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	server := s.ServerFactory.NewServer(server.HTTPServerProps{
		DB:   db,
		Port: config.Port,
	})

	return server.Start()
}
