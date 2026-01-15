package leakybucket

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucketRateLimiter struct {
	capacity     int
	water        int
	leakRate     float64
	lastLeakedAt time.Time
	mu           sync.Mutex
}

func NewLeakyBucketRateLimiter(capacity int, leakInterval time.Duration) *LeakyBucketRateLimiter {
	return &LeakyBucketRateLimiter{
		capacity:     capacity,
		water:        0,
		leakRate:     float64(capacity) / leakInterval.Seconds(),
		lastLeakedAt: time.Now(),
	}
}

func (rc *LeakyBucketRateLimiter) Leak() {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(rc.lastLeakedAt)
	leakedWater := int(elapsedTime.Seconds() * rc.leakRate)
	rc.water = max(0, rc.water-leakedWater)
	rc.lastLeakedAt = currentTime
	fmt.Printf("[Leak] Time: %v | Water leaked: %d | Current Water: %d\n",
		currentTime.Format("15:04:05.000"), leakedWater, rc.water)
}

func (rc *LeakyBucketRateLimiter) Allow() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.Leak()
	if rc.water < rc.capacity {
		rc.water += 1
		return true
	}
	return false
}

func (rc *LeakyBucketRateLimiter) String() string {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return fmt.Sprintf("LeakyBucket{capacity=%d, water=%d, leakRate=%.2f, lastLeakedAt=%v}",
		rc.capacity, rc.water, rc.leakRate, rc.lastLeakedAt,
	)
}
