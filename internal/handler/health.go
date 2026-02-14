package handler

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

// HealthResponse is the JSON structure returned by the health endpoint.
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Sentry    string `json:"sentry"`
}

// HandleHealth responds with the service health status including
// Sentry connectivity. Pass ?sentry_test=1 to send a test event.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	sentryStatus := "disabled"

	hub := sentry.GetHubFromContext(r.Context())
	if hub == nil {
		hub = sentry.CurrentHub()
	}

	if hub.Client() != nil {
		sentryStatus = "enabled"

		if r.URL.Query().Get("sentry_test") == "1" {
			hub.CaptureMessage("health check test event from docker-api-demo")
			sentryStatus = "test event sent"
		}
	}

	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Sentry:    sentryStatus,
	}

	writeJSON(w, http.StatusOK, resp)
}
