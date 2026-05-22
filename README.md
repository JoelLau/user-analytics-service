# User Analytics Service

## Pre-requisites

1. [docker](https://www.docker.com/products/docker-desktop/)
1. [docker compose](https://docs.docker.com/compose/)
1. [goose](https://github.com/pressly/goose)

## Instructions

1. create .env file `cp .env.example . env`
    (modify where required)
1. start external services (postgresql, migrations) `docker compose up -d`
1. install dependencies `go mod tidy`
1. generate database files (sqlc) and
    server interface (oapi-codegen) `go generate ./...`
