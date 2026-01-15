package tokenbucket

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucketRateLimiter struct {
	capacity       int
	tokens         int
	refillRate     float64
	lastRefilledAt time.Time
	mu             sync.Mutex
}

func NewTokenBucketRateLimiter(capacity int, refillInterval time.Duration) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     float64(capacity) / refillInterval.Seconds(),
		lastRefilledAt: time.Now(),
	}
}
func (rc *TokenBucketRateLimiter) Refill() {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(rc.lastRefilledAt)
	tokensToAdd := int(elapsedTime.Seconds() * rc.refillRate)
	rc.tokens = min(rc.capacity, rc.tokens+tokensToAdd)
	rc.lastRefilledAt = currentTime
	fmt.Printf("[Refill] Time: %v | Tokens added: %d | Current tokens: %d\n",
		currentTime.Format("15:04:05.000"), tokensToAdd, rc.tokens)
}
func (rc *TokenBucketRateLimiter) String() string {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return fmt.Sprintf("TokenBucket{capacity=%d, tokens=%d, refillRate=%.2f, lastRefilledAt=%v}",
		rc.capacity, rc.tokens, rc.refillRate, rc.lastRefilledAt)
}
func (rc *TokenBucketRateLimiter) Allow() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.Refill()
	if rc.tokens >= 1 {
		rc.tokens -= 1
		return true
	}
	return false
}
