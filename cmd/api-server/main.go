package main

import (
	"fmt"
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/routes"
	"github.com/milesjpool/kahd/internal/env"
)

func main() {
	http.HandleFunc("/", routes.NotFound)

	port := ":" + env.GetOrDefault("PORT", "8080")

	fmt.Println("Starting server on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
