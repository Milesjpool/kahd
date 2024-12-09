package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/milesjpool/kahd/cmd/api-server/internal/routes"
	"github.com/milesjpool/kahd/cmd/api-server/internal/routes/status"
	"github.com/milesjpool/kahd/internal/env"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	port := env.GetOrDefault("PORT", "8080")

	return startWebServer(port)
}

func startWebServer(port string) error {
	http.HandleFunc("/", routes.NotFound)
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		routes.Status(w, r, status.Context{
			Checks: map[string]func() bool{},
		})
	})

	fmt.Println("Starting server on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
