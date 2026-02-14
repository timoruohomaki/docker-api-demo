# CLAUDE.md — docker-api-demo

## Project Overview

A demo Go REST API designed to run in a Docker container behind an Nginx reverse proxy.
Part of a personal infrastructure project for learning Docker containerization and CI/CD.

Infrastructure context lives in the `backend01` repository (Nginx configs, SSL,
runbooks, deployment scripts). This app is deployed to `api.ruohomaki.fi` via
Nginx reverse proxy to `127.0.0.1:8080`.

## Architecture

- **Language:** Go 1.23+
- **Framework:** Standard library `net/http` with method-based routing (Go 1.22+ patterns)
- **Configuration:** Environment variables (12-factor app style)
- **Monitoring:** Sentry (developer subscription) — optional, enabled via SENTRY_DSN
- **Deployment:** Docker container behind host Nginx reverse proxy
- **Databases:** External PaaS services (Neon/Supabase for Postgres, MongoDB Atlas, etc.)

## Dependencies

- `github.com/getsentry/sentry-go` (v0.40.0) — error tracking and performance tracing
- `github.com/getsentry/sentry-go/http` — net/http middleware integration

No other external dependencies. Everything else uses Go's standard library.

## Project Structure

```
docker-api-demo/
├── cmd/server/main.go       # Entry point — wires config, Sentry, middleware, server
├── internal/
│   ├── config/config.go     # Environment-based configuration with validation
│   ├── handler/
│   │   ├── routes.go        # Route registration
│   │   ├── health.go        # GET /health — status + Sentry check + optional test event
│   │   ├── hello.go         # GET /api/hello — demo greeting endpoint
│   │   ├── response.go      # JSON response helpers (writeJSON, writeError)
│   │   └── handler_test.go  # Unit tests for health and hello handlers
│   ├── middleware/
│   │   ├── logging.go       # Request logging (method, path, status, duration)
│   │   └── sentry.go        # Sentry panic recovery and performance tracing
│   ├── monitoring/
│   │   └── sentry.go        # Sentry SDK initialization and cleanup
│   └── server/
│       └── server.go        # HTTP server creation and graceful shutdown
├── Dockerfile               # Multi-stage: golang:1.23-alpine → alpine:3.21
├── docker-compose.yml       # Local dev (builds locally, binds 127.0.0.1:8080)
├── docker-compose.prod.yml  # Production (pulls from ghcr.io)
├── .github/workflows/
│   ├── ci.yml               # Test + build on push/PR to main
│   └── cd.yml               # Build image → push to ghcr.io → deploy via SSH
├── .dockerignore
├── .env                     # Local env vars (not committed)
├── go.mod
└── go.sum
```

## Conventions

- Files should not exceed ~100 lines of code
- Production-grade error handling (no silent failures)
- Graceful shutdown on SIGINT/SIGTERM with 10-second drain timeout
- Configuration via environment variables, never hardcoded
- Dates and times in ISO 8601 / RFC 3339 format
- Structured log output with timestamps
- Container ports bind to 127.0.0.1 (Docker bypasses UFW on 0.0.0.0)
- Use `docker compose` (v2, space not hyphen)
- Sentry is optional — runs fine without SENTRY_DSN set

## Build & Run

### Local development (without Docker)

```bash
go run ./cmd/server
```

Note: Go does not read `.env` files. Export variables manually or use Docker Compose.

### Docker

```bash
docker compose up --build
```

Docker Compose reads `.env` automatically and passes variables to the container.

### Test

```bash
go test ./...
```

## Environment Variables

| Variable            | Default       | Description                    |
|---------------------|---------------|--------------------------------|
| PORT                | 8080          | HTTP listen port               |
| HOST                | 0.0.0.0       | HTTP listen address            |
| LOG_LEVEL           | info          | Logging verbosity              |
| SENTRY_DSN          | (empty)       | Sentry DSN — empty = disabled  |
| SENTRY_ENVIRONMENT  | development   | Sentry environment tag         |
| APP_VERSION         | 0.1.0         | Release version sent to Sentry |

## API Endpoints

| Method | Path                  | Description                           |
|--------|-----------------------|---------------------------------------|
| GET    | /health               | Health check + Sentry status          |
| GET    | /health?sentry_test=1 | Health check + send Sentry test event |
| GET    | /api/hello            | Demo greeting (accepts ?name=)        |

## Middleware Chain

Requests pass through: Sentry (panic recovery + tracing) → Request Logger → Handler

## CI/CD Pipeline

**CI** (`.github/workflows/ci.yml`) — runs on every push and PR to `main`:
- Downloads Go dependencies
- Runs `go test -v -race ./...`
- Verifies the binary compiles

**CD** (`.github/workflows/cd.yml`) — runs on push to `main` only:
- Builds Docker image
- Pushes to `ghcr.io/timoruohomaki/docker-api-demo` (tagged with commit SHA + `latest`)
- SSHs into server as `deploy`, pulls the new image, restarts the container

**Required GitHub Secrets:** `SERVER_HOST`, `SERVER_USER`, `SERVER_SSH_KEY`, `SERVER_PORT`
**Required GitHub Environment:** `production`

## Related Repositories

- **backend01** — Server infrastructure: Nginx configs, SSL snippets, deployment
  runbooks, static sites. The `api.ruohomaki.fi` Nginx config lives there.

## Current Status

- [x] Go API with clean architecture
- [x] Docker multi-stage build (non-root, ~15 MB image)
- [x] Sentry error tracking and performance tracing
- [x] Health endpoint with Sentry status
- [x] Graceful shutdown
- [x] Unit tests for handlers
- [x] GitHub Actions CI pipeline (tests + build on push/PR)
- [x] GitHub Actions CD pipeline (image push to ghcr.io + SSH deploy)
- [x] Server-side setup: SSH deploy key, GitHub Secrets, production compose
