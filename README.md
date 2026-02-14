# docker-api-demo

A minimal Go REST API running in a Docker container — built as a learning project for containerized deployment behind Nginx.

## Quick Start

### Run locally (Go required)

```bash
go run ./cmd/server
```

### Run with Docker

```bash
docker compose up --build
```

### Test the endpoints

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/hello
curl http://localhost:8080/api/hello?name=Timo
```

## API Endpoints

| Method | Path            | Description             |
|--------|-----------------|-------------------------|
| GET    | `/health`       | Health check            |
| GET    | `/api/hello`    | Greeting (accepts `?name=`) |

## Project Structure

```
cmd/server/          → Application entry point
internal/config/     → Environment-based configuration
internal/handler/    → HTTP request handlers
internal/middleware/ → Request logging
internal/server/     → HTTP server lifecycle
```

## Configuration

All configuration is via environment variables. See `.env` for defaults.

## Docker

The Dockerfile uses a multi-stage build: compiles in `golang:1.23-alpine`, runs in `alpine:3.21` as a non-root user. The final image is ~15 MB.
