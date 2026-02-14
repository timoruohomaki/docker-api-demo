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
- **Deployment:** Docker container behind host Nginx reverse proxy
- **Databases:** External PaaS services (Neon/Supabase for Postgres, MongoDB Atlas, etc.)

## Project Structure

```
docker-api-demo/
├── cmd/server/          # Application entry point
├── internal/
│   ├── config/          # Environment-based configuration
│   ├── handler/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware (logging, etc.)
│   └── server/          # HTTP server setup and lifecycle
├── Dockerfile           # Multi-stage production build
├── docker-compose.yml   # Local development setup
└── .env                 # Local env vars (not committed)
```

## Conventions

- Files should not exceed ~100 lines of code
- Production-grade error handling (no silent failures)
- Graceful shutdown on SIGINT/SIGTERM
- Configuration via environment variables, never hardcoded
- Dates and times in ISO 8601 / RFC 3339 format
- Structured log output with timestamps
- Container ports bind to 127.0.0.1 (Docker bypasses UFW on 0.0.0.0)
- Use `docker compose` (v2, space not hyphen)

## Build & Run

### Local development (without Docker)

```bash
go run ./cmd/server
```

### Docker

```bash
docker compose up --build
```

### Test

```bash
go test ./...
```

## Environment Variables

| Variable   | Default       | Description            |
|------------|---------------|------------------------|
| PORT       | 8080          | HTTP listen port       |
| HOST       | 0.0.0.0       | HTTP listen address    |
| LOG_LEVEL  | info          | Logging verbosity      |

## API Endpoints

| Method | Path       | Description              |
|--------|------------|--------------------------|
| GET    | /health    | Health check (readiness)  |
| GET    | /api/hello | Demo greeting endpoint    |
