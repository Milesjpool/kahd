package routes

import (
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/response"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response.WriteJson(map[string]string{}, w)
}
