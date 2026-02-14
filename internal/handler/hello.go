package handler

import (
	"net/http"
	"time"
)

// HelloResponse is the JSON structure returned by the hello endpoint.
type HelloResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// HandleHello responds with a greeting message.
func HandleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	resp := HelloResponse{
		Message:   "Hello, " + name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusOK, resp)
}
