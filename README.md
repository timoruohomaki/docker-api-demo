# docker-api-demo

A minimal Go REST API running in a Docker container — built as a learning project for containerized deployment behind Nginx.

Deployed to `api.ruohomaki.fi` via Nginx reverse proxy. Infrastructure configuration lives in the [backend01](https://github.com/timoruohomaki/backend01) repository.

## Quick Start

### Run locally (Go required)

```bash
go run ./cmd/server
```

Note: Go does not read `.env` files automatically. Either export variables in your shell or use Docker Compose.

### Run with Docker

```bash
docker compose up --build
```

Docker Compose reads the `.env` file and passes variables to the container.

### Run tests

```bash
go test ./...
```

### Test the endpoints

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/hello
curl http://localhost:8080/api/hello?name=Timo
curl "http://localhost:8080/health?sentry_test=1"
```

## API Endpoints

| Method | Path                    | Description                           |
|--------|-------------------------|---------------------------------------|
| GET    | `/health`               | Health check + Sentry status          |
| GET    | `/health?sentry_test=1` | Health check + send Sentry test event |
| GET    | `/api/hello`            | Greeting (accepts `?name=`)           |

## Project Structure

```
cmd/server/          → Entry point, wires config, middleware, server
internal/config/     → Environment-based configuration
internal/handler/    → HTTP request handlers + response helpers
internal/middleware/ → Request logging, Sentry panic recovery
internal/monitoring/ → Sentry SDK initialization
internal/server/     → HTTP server lifecycle + graceful shutdown
```

## Configuration

All configuration is via environment variables. See `.env` for defaults.

| Variable            | Default       | Description                    |
|---------------------|---------------|--------------------------------|
| PORT                | 8080          | HTTP listen port               |
| HOST                | 0.0.0.0       | HTTP listen address            |
| LOG_LEVEL           | info          | Logging verbosity              |
| SENTRY_DSN          | (empty)       | Sentry DSN — empty = disabled  |
| SENTRY_ENVIRONMENT  | development   | Sentry environment tag         |
| APP_VERSION         | 0.1.0         | Release version sent to Sentry |

## Monitoring

Sentry integration is optional. Without `SENTRY_DSN` set, the app runs normally and the `/health` endpoint reports `"sentry": "disabled"`. When enabled, the middleware captures panics and creates performance traces for all requests.

## Docker

The Dockerfile uses a multi-stage build: compiles in `golang:1.23-alpine`, runs in `alpine:3.21` as a non-root user. The final image is approximately 15 MB. Container ports bind to `127.0.0.1` to prevent Docker from bypassing the host firewall.
