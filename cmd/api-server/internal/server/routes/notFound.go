package routes

import (
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/server/response"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	response.WriteJson(map[string]string{"error": "not found"}, w)
}
