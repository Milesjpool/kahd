package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/resource" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"resource": "123"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	t.Run("Get_ValidResource", func(t *testing.T) {
		client := &Client{URL: ts.URL}

		resp, err := client.Get("resource")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()
		content, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"resource": "123"}`, string(content))
	})

	t.Run("Get_NotFound", func(t *testing.T) {
		client := &Client{URL: ts.URL}

		resp, err := client.Get("unknown")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
