# User Analytics Service

## Pre-requisites

1. [docker](https://www.docker.com/products/docker-desktop/)
1. [docker compose](https://docs.docker.com/compose/)
1. [Go](https://go.dev/) >= 1.26.3
1. [optional] [goose](https://github.com/pressly/goose)
1. [optional] [sqlc](https://sqlc.dev/)

## Running this project

these are the recommended instructions for running this project
on a linux-based machine (e.g. macos)

1. configure project by setting environment variables
    1. `cp .env.example .env`: create an _.env_ file using the provided example
    1. _modify the contents of .env_, PLEASE CHANGE THE DEFAULT PASSWORD
    1. `export $(grep -v '^#' .env | xargs)`: set the contents of the .env file
        into your local environment
1. set up database (run goose migrations, seed demo data)
    1. `docker compose up -d` run docker compose in background
        - spins up postgresql database instance
        - runs database migrations to create necessary tables, etc
        - seeds database with data for this demo
1. install go dependencies
    - `go mod tidy`
1. start the user analytics service:
    - `go run cmd/web-api/main.go` starts up the REST API
1. hit the endpoints you need
    1. [swaggerdocs are available at /api/swagger/]([http://localhost:8080/api/swagger/])
    1. click the "Try it out" button on the corresponding endpoint
    1. modify the parameters
    1. click execute to see results
    - NOTE: demo data ranges between `2026-03-23` (inclusive) and `2026-05-22` (inclusive)

## Endpoints

check [openapi.yaml] for source of truth

```plaintext
GET /api/v1/analytics/users/daily/{day}
```

```plaintext
GET /api/v1/analytics/users/monthly/{month}
```

## Commands

| Commands | Remarks |
| --- | --- |
| `docker compose up -d` | starts external services |
| `docker compose down -v` | stops external services |
| `go mod tidy` | install go dependencies |
| `cp .env.example .env` | create .env from examples |
| `export $(grep -v '^#' .env | xargs)` | load environment variables |
| `go run cmd/web-api/main.go` | starts user analytics service REST API server |
| `go test ./...` | run tests (requires docker) |
| `go generate ./...` | code generation - sql helpers, openapi interfaces |

## Project Layout

```plaintext
cmd/web-api/        entrypoint — wires config, database pool, queries, clock, and HTTP handler
config/             environment-based configuration and logger setup
docs/               requirements, tech stack, and architecture documents
migrations/         goose SQL migrations (source of truth for the database schema)
queries/            sqlc SQL source files and the Go code generated from them
seed/               demo data as plain INSERT statements (run with goose --no-versioning)
server/             HTTP handler, server implementation, oapi-codegen generated types, and tests
openapi.yaml        OpenAPI 3.0 spec — served at runtime under /api/openapi.yaml
sqlc.yaml           sqlc code generation config
docker-compose.yaml spins up postgres, runs migrations, and seeds demo data
```
