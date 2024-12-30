package test_suites

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/milesjpool/kahd/e2e-tests/internal"
	"github.com/milesjpool/kahd/e2e-tests/internal/assertions"
	"github.com/milesjpool/kahd/internal/model"
)

type WebApiTestSuite struct {
	internal.TestSuite
	URL string
}

func (s *WebApiTestSuite) Run(t *internal.TestContext) {
	restClient := resty.New()

	t.Run("it gets 404 for an unknown resource", func(t *internal.TestContext) {
		resp, err := restClient.R().Get(s.URL + "/unknown")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusNotFound, resp.StatusCode())
	})

	t.Run("it retrieves API status", func(t *internal.TestContext) {
		status := &model.Status{}
		resp, err := restClient.R().SetResult(status).Get(s.URL + "/status")

		assertions.NoErr(t, err, "error making request: %v", err)
		assertions.Equals(t, http.StatusOK, resp.StatusCode())
		assertions.Equals(t, "application/json", resp.Header().Get("Content-Type"))
		assertions.Equals(t, "healthy", status.DatabaseConnection)
	})
}
