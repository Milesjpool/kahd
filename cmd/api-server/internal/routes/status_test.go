package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/milesjpool/kahd/cmd/api-server/internal/routes/status"
	"github.com/stretchr/testify/assert"
)

func TestStatusWithNoChecks(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Status(rr, req, status.Context{})

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{}`, rr.Body.String())
}

func TestStatusWithSetOfChecks(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Status(rr, req, status.Context{
		Checks: map[string]func() bool{
			"database_connection": func() bool {
				return true
			},
			"another_check": func() bool {
				return false
			},
		},
	})

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"database_connection":"healthy","another_check":"unhealthy"}`, rr.Body.String())
}
