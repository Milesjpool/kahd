package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/milesjpool/kahd/cmd/api-server/internal/server/routes/status"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {

	t.Run("it returns a 200 status when there are no checks", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)

		Status(rr, req, status.Context{})

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, `{}`, rr.Body.String())
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	})

	t.Run("it returns a 200 status when there are checks", func(t *testing.T) {
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
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	})
}
