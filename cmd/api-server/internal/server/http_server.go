package server

import (
	"fmt"
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/logging"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server/routes"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server/routes/status"
)

type HTTPServerFactory struct {
	Logger logging.Logger
}

type HTTPServerProps struct {
	DB   database.Database
	Port string
}

func (f *HTTPServerFactory) NewServer(props HTTPServerProps) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.NotFound)
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		routes.Status(w, r, status.Context{
			Checks: map[string]func() bool{
				"database_connection": func() bool {
					return props.DB.Ping() == nil
				},
			},
		})
	})

	return &httpServerAdapter{
		serverListener: &http.Server{
			Addr:    ":" + props.Port,
			Handler: mux,
		},
		beforeStart: func() {
			f.Logger.Info("Starting server at: " + props.Port)
		},
	}
}

type serverListener interface {
	ListenAndServe() error
}

type httpServerAdapter struct {
	serverListener
	beforeStart func()
}

func (s *httpServerAdapter) Start() error {
	s.beforeStart()

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
