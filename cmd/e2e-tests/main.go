package main

import (
	"net/http"
	"os"

	"e2e-tests/internal"
	"e2e-tests/internal/assertions"

	"github.com/milesjpool/kahd/pkg/api"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost:8080"
	}

	t := &internal.TestContext{}
	apiClient := &api.Client{URL: "http://" + host}

	t.Run("it gets 404 for an unknown resource", func(t *internal.TestContext) {
		resp, err := apiClient.Get("unknown")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it retrieves API status", func(t *internal.TestContext) {
		resp, err := apiClient.Get("status")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusOK, resp.StatusCode)
	})

	t.Close()
}
