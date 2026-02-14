package handler

import "net/http"

// NewRouter sets up all API routes and returns the configured mux.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HandleHealth)
	mux.HandleFunc("GET /api/hello", HandleHello)

	return mux
}
