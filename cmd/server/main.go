package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/timoruohomaki/docker-api-demo/internal/config"
	"github.com/timoruohomaki/docker-api-demo/internal/handler"
	"github.com/timoruohomaki/docker-api-demo/internal/middleware"
	"github.com/timoruohomaki/docker-api-demo/internal/monitoring"
	"github.com/timoruohomaki/docker-api-demo/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	sentryCleanup, err := monitoring.InitSentry(cfg)
	if err != nil {
		log.Fatalf("failed to initialize sentry: %v", err)
	}
	defer sentryCleanup()

	mux := handler.NewRouter()

	// Middleware chain: Sentry (outer) → Logger → Handler
	wrapped := middleware.SentryHandler(middleware.RequestLogger(mux))

	srv := server.New(cfg, wrapped)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := server.Run(ctx, srv); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
