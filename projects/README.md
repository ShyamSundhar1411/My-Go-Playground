# Projects

Applied, end-to-end applications that build on the concepts explored
in `concepts/`. Where a concept module is a single idea in isolation,
a project here is a small, production-shaped service.

Each project:

- lives in its own directory
- has its own `go.mod`, kept independent so its dependencies don't leak into `concepts/`, and is added to the root `go.work`
- has its own README describing what it does and how to run it

## Current projects

- [`rest-api`](./rest-api) — a REST API scaffolded with a standard layered Go structure (`cmd`, `internal/{handler,service,repository,model,middleware,config}`)

## Adding a new project

1. `mkdir projects/<name> && cd projects/<name> && go mod init github.com/ShyamSundhar1411/<name>`
2. Add `./projects/<name>` to the root `go.work`
3. Write a README describing the project's purpose and structure
