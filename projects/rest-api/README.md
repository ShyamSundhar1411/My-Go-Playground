# REST API

A REST API built on top of the concepts explored in this repo (rate
limiting, concurrency, HTTP internals), structured the way a
production Go service would be laid out.

## Layout

- `cmd/api` — entrypoint; wires config, router, and starts the server
- `internal/config` — environment-based configuration
- `internal/router` — route registration and middleware wiring
- `internal/handler` — HTTP handlers (thin: parse request, call service, write response)
- `internal/service` — business logic
- `internal/repository` — data access (database, external APIs)
- `internal/model` — domain types and DTOs
- `internal/middleware` — shared HTTP middleware (logging, auth, rate limiting, ...)
- `migrations/` — SQL migration files, once a database is introduced

This module has its own `go.mod` so its dependencies (router, DB
driver, etc.) don't leak into the concept modules under `concepts/`.
It's tied into the repo via the root `go.work`, so `go build ./...`
and editor tooling still work across both.

## Running

```bash
cd projects/rest-api
go run ./cmd/api
```

Then:

```bash
curl localhost:8080/healthz
```

## Roadmap

- [ ] Add a real resource (users) end to end: model -> repository -> service -> handler
- [ ] Add request validation
- [ ] Add a database (likely Postgres) and wire up `migrations/`
- [ ] Reuse the rate limiters from `concepts/rate-limiting` as middleware
- [ ] Add integration tests
