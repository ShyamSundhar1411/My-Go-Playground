# Token Bucket Algorithm

The token bucket algorithm is a widely used rate limiting technique that controls the rate at which requests are processed.

Tokens are added to the bucket at a fixed rate. Each incoming request consumes one or more tokens. A request is allowed only if there are enough tokens available in the bucket. Unused tokens accumulate up to the maximum capacity of the bucke, allowing short bursts of traffic.

## Key Characteristics and Functions

- **Refill Rate:** Tokens are added at a fixed rate over time.
- **Capacity:** Maximum number of tokens the bucket can hold.
- **Allow():** Determines whether a request can proceed based on available tokens. If enough tokens are available, one or more tokens are consumed and the request is allowed; otherwise, it is denied.
- **Refill():** Adds tokens to the bucket according to the refill rate, up to the maximum capacity.

## Tradeoffs

- Requires careful handling of time and concurrency.
- Burstiness depends on bucket capacity.
- Not ideal for all scenarios—sometimes combined with other rate limiting strategies.