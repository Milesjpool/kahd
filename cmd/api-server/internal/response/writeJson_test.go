package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJson(t *testing.T) {
	t.Run("empty body", func(t *testing.T) {
		rr := httptest.NewRecorder()
		WriteJson(map[string]string{}, rr)

		assert.Equal(t, http.StatusOK, rr.Code, "expected status code to be equal")
		assert.JSONEq(t, `{}`, rr.Body.String(), "expected body to be equal")
	})

	t.Run("valid body", func(t *testing.T) {
		rr := httptest.NewRecorder()
		WriteJson(map[string]string{"key": "value"}, rr)

		assert.Equal(t, http.StatusOK, rr.Code, "expected status code to be equal")
		assert.JSONEq(t, `{"key":"value"}`, rr.Body.String(), "expected body to be equal")
	})
}
