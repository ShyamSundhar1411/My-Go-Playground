# Concepts

From-scratch, self-contained implementations of individual systems
concepts. Each folder here is a single idea: theory in a README, an
idiomatic Go implementation, and tests where it makes sense. These
stay small and dependency-free — bigger, applied work lives under
`projects/`.

## Index

- [`concurrency`](./concurrency) — goroutines, channels, sync primitives, deadlock conditions
- [`rate-limiting`](./rate-limiting) — fixed window, sliding window log, token bucket, leaky bucket
- [`networking`](./networking) — HTTP server/client internals, REST serialization

## Adding a new concept

1. `mkdir concepts/<topic>`
2. Add a README with the problem statement, theory, and tradeoffs
3. Add the implementation plus tests and a small `example/main.go` if it helps to see it run
