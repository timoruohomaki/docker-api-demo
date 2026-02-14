package middleware

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

// SentryHandler wraps the given handler with Sentry panic recovery
// and performance tracing. If Sentry is not initialized, the handler
// is returned unwrapped.
func SentryHandler(next http.Handler) http.Handler {
	if sentry.CurrentHub().Client() == nil {
		return next
	}

	handler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	return handler.Handle(next)
}
