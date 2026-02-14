package monitoring

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/timoruohomaki/docker-api-demo/internal/config"
)

const flushTimeout = 2 * time.Second

// InitSentry sets up the Sentry SDK. Returns a cleanup function that
// should be deferred in main. If DSN is empty, Sentry is skipped.
func InitSentry(cfg *config.Config) (func(), error) {
	if !cfg.SentryEnabled() {
		log.Println("sentry: disabled (no SENTRY_DSN set)")
		return func() {}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		Environment:      cfg.SentryEnv,
		Release:          cfg.AppVersion,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return nil, fmt.Errorf("sentry init failed: %w", err)
	}

	log.Printf("sentry: enabled (env=%s, release=%s)", cfg.SentryEnv, cfg.AppVersion)

	cleanup := func() {
		sentry.Flush(flushTimeout)
	}

	return cleanup, nil
}
