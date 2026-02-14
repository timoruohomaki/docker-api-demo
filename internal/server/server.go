package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/timoruohomaki/docker-api-demo/internal/config"
)

const (
	readTimeout     = 5 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 30 * time.Second
	shutdownTimeout = 10 * time.Second
)

// New creates a configured HTTP server ready to run.
func New(cfg *config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.ListenAddr(),
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

// Run starts the HTTP server and blocks until the context is cancelled,
// then performs a graceful shutdown.
func Run(ctx context.Context, srv *http.Server) error {
	errCh := make(chan error, 1)

	go func() {
		log.Printf("server is listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen error: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Println("shutdown signal received, draining connections...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	log.Println("server stopped gracefully")
	return nil
}
