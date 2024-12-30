package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDatabase struct {
	pingErr error
}

func (m *mockDatabase) Ping() error {
	return m.pingErr
}

func (m *mockDatabase) Close() error {
	return nil
}

type mockLogger struct {
	InfoCalls []string
}

func (m *mockLogger) Info(message string) {
	m.InfoCalls = append(m.InfoCalls, message)
}

func TestHTTPServerFactory_NewServer(t *testing.T) {
	t.Run("it configures the server with the correct port", func(t *testing.T) {
		factory := &HTTPServerFactory{}
		expectedPort := "8080"

		server := factory.NewServer(HTTPServerProps{
			DB:   &mockDatabase{},
			Port: expectedPort,
		})

		httpServer := server.(*httpServerAdapter).serverListener.(*http.Server)

		assert.Equal(t, ":"+expectedPort, httpServer.Addr)
	})

	t.Run("it returns 404 for unknown routes", func(t *testing.T) {
		factory := &HTTPServerFactory{}
		server := factory.NewServer(HTTPServerProps{
			DB:   &mockDatabase{},
			Port: "8080",
		})

		httpServer := server.(*httpServerAdapter).serverListener.(*http.Server)

		req := httptest.NewRequest("GET", "/unknown", nil)
		rr := httptest.NewRecorder()

		httpServer.Handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, `{"error":"not found"}`, rr.Body.String())
	})

	t.Run("status endpoint returns healthy database health", func(t *testing.T) {
		factory := &HTTPServerFactory{}
		server := factory.NewServer(HTTPServerProps{
			DB:   &mockDatabase{},
			Port: "8080",
		})

		httpServer := server.(*httpServerAdapter).serverListener.(*http.Server)

		req := httptest.NewRequest("GET", "/status", nil)
		rr := httptest.NewRecorder()

		httpServer.Handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"database_connection":"healthy"}`, rr.Body.String())
	})

	t.Run("status endpoint returns unhealthy database status", func(t *testing.T) {
		factory := &HTTPServerFactory{}
		server := factory.NewServer(HTTPServerProps{
			DB:   &mockDatabase{pingErr: errors.New("connection failed")},
			Port: "8080",
		})

		httpServer := server.(*httpServerAdapter).serverListener.(*http.Server)

		req := httptest.NewRequest("GET", "/status", nil)
		rr := httptest.NewRecorder()

		httpServer.Handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{"database_connection":"unhealthy"}`, rr.Body.String())
	})

	t.Run("it logs the server starting in the beforeStart function", func(t *testing.T) {
		logger := &mockLogger{}

		factory := &HTTPServerFactory{Logger: logger}
		server := factory.NewServer(HTTPServerProps{
			DB:   &mockDatabase{pingErr: errors.New("connection failed")},
			Port: "8080",
		})

		server.(*httpServerAdapter).beforeStart()

		assert.Contains(t, logger.InfoCalls, "Starting server at: 8080")
	})
}

type MockServer struct {
	listenAndServeCalled bool
	err                  error
}

func (m *MockServer) ListenAndServe() error {
	m.listenAndServeCalled = true
	return m.err
}

func TestHttpServerAdapter_Start(t *testing.T) {
	t.Run("it calls the beforeStart function if provided", func(t *testing.T) {
		beforeStartCalled := false
		adapter := &httpServerAdapter{serverListener: &MockServer{}, beforeStart: func() {
			beforeStartCalled = true
		}}
		err := adapter.Start()

		assert.NoError(t, err)
		assert.True(t, beforeStartCalled)
	})

	t.Run("it starts the server", func(t *testing.T) {
		mockServer := &MockServer{}
		adapter := &httpServerAdapter{serverListener: mockServer, beforeStart: func() {}}
		err := adapter.Start()

		assert.NoError(t, err)
		assert.True(t, mockServer.listenAndServeCalled)
	})

	t.Run("it returns an error if the server fails to start", func(t *testing.T) {
		expectedErr := errors.New("failed to start server")
		mockServer := &MockServer{err: expectedErr}
		adapter := &httpServerAdapter{serverListener: mockServer, beforeStart: func() {}}
		err := adapter.Start()

		assert.ErrorIs(t, err, expectedErr)
	})
}
