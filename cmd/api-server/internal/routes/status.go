package routes

import (
	"net/http"

	"github.com/milesjpool/kahd/cmd/api-server/internal/response"
	"github.com/milesjpool/kahd/cmd/api-server/internal/routes/status"
)

func Status(w http.ResponseWriter, r *http.Request, ctx status.Context) {
	status := map[string]string{}

	for key, check := range ctx.Checks {
		status[key] = "healthy"
		if !check() {
			status[key] = "unhealthy"
		}
	}

	w.WriteHeader(http.StatusOK)
	response.WriteJson(status, w)
}
