package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(
	body map[string]string,
	w http.ResponseWriter,
) {
	if jsonBody, err := json.Marshal(body); err == nil {
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonBody)
	}
}
