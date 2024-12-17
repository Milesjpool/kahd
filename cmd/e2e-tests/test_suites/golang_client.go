package test_suites

import (
	"github.com/milesjpool/kahd/e2e-tests/internal"
	"github.com/milesjpool/kahd/e2e-tests/internal/assertions"
	"github.com/milesjpool/kahd/pkg/api"
)

type GolangClientTestSuite struct {
	internal.TestSuite
	Host string
}

func (s *GolangClientTestSuite) Run(t *internal.TestContext) {
	apiClient := api.NewClient("http://" + s.Host)

	t.Run("it retrieves API status", func(t *internal.TestContext) {
		_, err := apiClient.Status()

		assertions.NoErr(t, err, "error making request: %v", err)
	})
}
