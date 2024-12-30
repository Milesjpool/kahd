package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	t.Run("it returns a 404 status", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)

		NotFound(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.JSONEq(t, `{"error":"not found"}`, rr.Body.String())
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	})
}
