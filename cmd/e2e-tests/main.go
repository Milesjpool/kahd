package main

import (
	"fmt"

	"github.com/milesjpool/kahd/e2e-tests/internal"
	"github.com/milesjpool/kahd/e2e-tests/test_suites"
	"github.com/milesjpool/kahd/internal/env"
)

func main() {
	host := env.GetOrDefault("HOST", "localhost:8080")

	t := &internal.TestContext{}

	t.Init(func(t *internal.TestContext) {
		internal.WaitForServer(t, host, 5)
	})
	defer t.Close()

	suites := []internal.TestSuite{
		&test_suites.WebApiTestSuite{URL: fmt.Sprintf("http://%s", host)},
	}

	for _, suite := range suites {
		suite.Run(t)
	}
}
