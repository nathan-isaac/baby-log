# Project Baby Log

Project Baby Log is a simple web application to keep track of baby activities.

## Getting Started


## Environment

Create a .env file in the root of the project with the following content

```bash
PORT=8080
HOST=localhost
DATABASE_URL=http://localhost:8080

ADMIN_USER=admin
ADMIN_PASSWORD=admin

SENTRY_DSN="https://{}.ingest.us.sentry.io/{}"
SENTRY_SERVER_NAME=local
```

## Database

```bash
GOOSE_MIGRATION_DIR=schema GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=.storage/local.db goose up
```

## Prerequisites

- https://github.com/cosmtrek/air
- https://github.com/pressly/goose
- https://docs.sqlc.dev/
- https://github.com/pvolok/mprocs

```bash
go install github.com/cosmtrek/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/a-h/templ/cmd/templ@latest
```
