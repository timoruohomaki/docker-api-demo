package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse is the JSON structure returned by the health endpoint.
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// HandleHealth responds with the service health status.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusOK, resp)
}
