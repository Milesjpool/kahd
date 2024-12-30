package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server"
	"github.com/stretchr/testify/assert"
)

type mockConfigLoader struct {
	err          error
	port         string
	dbConnection string
}

func (m *mockConfigLoader) Load() (APIConfig, error) {
	port := m.port
	if port == "" {
		port = "8080"
	}
	dbConnection := m.dbConnection
	if dbConnection == "" {
		dbConnection = "my-db-connection-string"
	}
	return APIConfig{
		Port:         port,
		DBConnection: dbConnection,
	}, m.err
}

type mockDatabase struct {
	dbConnection string
	closed       bool
}

func (m *mockDatabase) Close() error {
	m.closed = true
	return nil
}

func (m *mockDatabase) Ping() error {
	return nil
}

type mockDatabaseConnector struct {
	err error
	db  *mockDatabase
}

func (m *mockDatabaseConnector) Connect(dbConnection string) (database.Database, error) {
	db := m.db
	if db == nil {
		db = &mockDatabase{}
	}
	db.dbConnection = dbConnection
	return db, m.err
}

type mockServer struct {
	err  error
	port string
}

func (m *mockServer) Start() error {
	return m.err
}

type mockServerFactory struct {
	server *mockServer
}

func (m *mockServerFactory) NewServer(props server.HTTPServerProps) server.Server {
	server := m.server
	if server == nil {
		server = &mockServer{}
	}
	server.port = props.Port
	return server
}

func TestAPIServer_Start(t *testing.T) {
	t.Run("it starts the server", func(t *testing.T) {
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{},
			DatabaseConnector: &mockDatabaseConnector{},
			ServerFactory:     &mockServerFactory{},
		}
		err := server.Start()
		assert.NoError(t, err)
	})

	t.Run("it returns an error if the config loader fails", func(t *testing.T) {
		expectedErr := errors.New("failed to load config")
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{err: expectedErr},
			DatabaseConnector: &mockDatabaseConnector{},
			ServerFactory:     &mockServerFactory{},
		}
		err := server.Start()
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("it closes the database connection", func(t *testing.T) {
		db := &mockDatabase{}
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{},
			DatabaseConnector: &mockDatabaseConnector{db: db},
			ServerFactory:     &mockServerFactory{},
		}
		assert.False(t, db.closed)
		server.Start()
		assert.True(t, db.closed)
	})

	t.Run("it connects to the database with the loaded config", func(t *testing.T) {
		expectedDBConnection := fmt.Sprintf("mock-connection-%d", rand.Intn(65535))
		db := &mockDatabase{}
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{dbConnection: expectedDBConnection},
			DatabaseConnector: &mockDatabaseConnector{db: db},
			ServerFactory:     &mockServerFactory{},
		}
		server.Start()
		assert.Equal(t, expectedDBConnection, db.dbConnection)
	})

	t.Run("it returns an error if the database connector fails", func(t *testing.T) {
		expectedErr := errors.New("failed to connect to database")
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{},
			DatabaseConnector: &mockDatabaseConnector{err: expectedErr},
			ServerFactory:     &mockServerFactory{},
		}
		err := server.Start()
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("it starts the server with the loaded config", func(t *testing.T) {
		expectedPort := fmt.Sprintf("%d", rand.Intn(65535))
		server := &mockServer{}
		unit := &APIServer{
			ConfigLoader:      &mockConfigLoader{port: expectedPort},
			DatabaseConnector: &mockDatabaseConnector{},
			ServerFactory:     &mockServerFactory{server: server},
		}
		unit.Start()
		assert.Equal(t, expectedPort, server.port)
	})

	t.Run("it returns an error if the server fails to start", func(t *testing.T) {
		expectedErr := errors.New("failed to start server")
		server := &APIServer{
			ConfigLoader:      &mockConfigLoader{},
			DatabaseConnector: &mockDatabaseConnector{},
			ServerFactory:     &mockServerFactory{server: &mockServer{err: expectedErr}},
		}
		err := server.Start()
		assert.ErrorIs(t, err, expectedErr)
	})
}
