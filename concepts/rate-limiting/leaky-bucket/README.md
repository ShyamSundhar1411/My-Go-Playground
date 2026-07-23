# Leaky Bucket Algorithm

The Leaky Bucket algorithm is a rate limiting technique that models incoming requests as water flowing into a bucket with a fixed capacity and a constant leak rate.

Requests are added to the bucket as they arrive, while the bucket leaks at a steady rate over time. If the incoming request rate exceeds the leak rate and the bucket reaches its capacity, additional requests are delayed or rejected.

This approach smooths bursty traffic by enforcing a consistent processing rate.

## Key Characteristics and Functions
- **Capacity:** Maximum amoung of requests(water) the bucket can hold
- **Leak Rate:** Fixed rate at which requests are processed(leaked)
- **Allow():** Determines whether a request can be processed based on the bucket capacity and leak rate
- **Leak():** Removes requests from the bucket based on elapsed time

## Tradeoffs
- Can introduce latency
- Not ideal for short bursts
- Less flexible than token bucket