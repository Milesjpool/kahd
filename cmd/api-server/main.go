package main

import (
	"fmt"
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/routes"
	"github.com/milesjpool/kahd/internal/env"
)

func main() {
	http.HandleFunc("/", routes.NotFound)
	http.HandleFunc("/status", routes.Status)

	port := env.GetOrDefault("PORT", "8080")

	fmt.Println("Starting server on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
