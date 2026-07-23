# My Go Playground

This repository contains clean, from-scratch implementations of
core systems and backend engineering concepts, built using Go.

The goal is to develop strong systems intuition by implementing fundamental primitives and well-known system patterns from first principles, with an emphasis on correctness, concurrency, and real-world behavior

## Focus areas
- Low-level fundamentals (atomics, locks, goroutines, channels)
- Concurrency problems and deadlock scenarios
- Rate limiting algorithms (fixed window, sliding window, token bucket, leaky bucket)
- Caching, consistency, and state management
- Fault tolerance patterns (retries, backoff, circuit breakers)
- Backpressure, scheduling, and time-based coordination

## Repository Structure

The repo is split into two top-level areas:

- **[`concepts/`](./concepts)** — one idea per folder (concurrency, rate-limiting, networking). Each includes a clear problem statement and theory, an idiomatic Go implementation, concurrency and performance considerations, edge cases and tradeoffs, and tests/examples where applicable. These stay small and dependency-free.
- **[`projects/`](./projects)** — applied, end-to-end services that build on those concepts (e.g. [`rest-api`](./projects/rest-api)), laid out the way a production Go service would be. Each project has its own `go.mod` so its dependencies don't leak into `concepts/`, tied together with the root `go.work`.

This repository is both a personal learning playground and an
open-source reference for engineers who want to build strong
systems intuition using Go.