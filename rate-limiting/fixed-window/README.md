# Fixed Window Algorithm

The Fixed Window algorithm is a rate limiting technique that allows a maximum number of requests within a fixed time window (for example, one minute or one hour).

Each incoming request increments a counter associated with the current window. Once the counter reaches the configured limit, further requests are rejected or delayed until the current window expires and a new window begins, at which point the counter is reset.

## Key Characteristics and Functions
- **Limit:** Maximum number of requests for the window frame
- **Count:** Current requests in the present window frame
- **WindowSize:** Duration of a window frame
- **Allow:** Determines to whether process or reject the request

## Tradeoffs
- Cannot handle boundary bursts
- Not suitable for high-scale systems
