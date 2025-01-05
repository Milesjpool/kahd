package main

import (
	"errors"
	"fmt"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/logging"
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
	Logger            logging.Logger
}

func (s *APIServer) Start() error {
	config, err := s.ConfigLoader.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	db, err := s.DatabaseConnector.Connect(config.DBConnection)
	if errors.Is(err, database.ErrDatabaseNotReachable) {
		s.Logger.Error("database not currently reachable. Error: %w", err)
	} else if err != nil {
		return fmt.Errorf("failed to initialize database connection: %w", err)
	}
	defer db.Close()

	server := s.ServerFactory.NewServer(server.HTTPServerProps{
		DB:   db,
		Port: config.Port,
	})

	return server.Start()
}
