# Sliding Window Log Algorithm

The Sliding Window Log algorithm is a rate limiting technique that enforces limits by storing the timestamps of individual requests within a rolling time window.

As new requests arrive, timestamps that fall outside the configured window size are removed (expired) from the log. If the number of remaining timestamps reaches the configured limit, additional requests are rejected or delayed until older entries expire.

## Key Characteristics and Functions
- **Limit:** Maximum number of requests for the window frame
- **WindowSize:** Duration of a window frame
- **Request Log:** Queue of timestamps representing recent requests
- **Allow:** Determines to whether process or reject the request

## Tradeoffs
- Memory intensive algorithm
- Not suitable for higher traffic
