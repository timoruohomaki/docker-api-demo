package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse is the standard JSON error structure.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// writeJSON serializes data as JSON and writes it to the response.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode JSON response: %v", err)
	}
}

// writeError sends a standardized JSON error response.
func writeError(w http.ResponseWriter, status int, message string) {
	resp := ErrorResponse{
		Error: message,
		Code:  status,
	}
	writeJSON(w, status, resp)
}
