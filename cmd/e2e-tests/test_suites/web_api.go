package test_suites

import (
	"net/http"

	"github.com/milesjpool/kahd/e2e-tests/internal"
	"github.com/milesjpool/kahd/e2e-tests/internal/assertions"
)

type WebApiTestSuite struct {
	internal.TestSuite
	URL string
}

func (s *WebApiTestSuite) Run(t *internal.TestContext) {
	httpClient := &http.Client{}

	t.Run("it gets 404 for an unknown resource", func(t *internal.TestContext) {
		resp, err := httpClient.Get(s.URL + "/unknown")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it retrieves API status", func(t *internal.TestContext) {
		resp, err := httpClient.Get(s.URL + "/status")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusOK, resp.StatusCode)
	})
}
