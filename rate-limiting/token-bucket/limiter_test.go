package tokenbucket

import (
	"testing"
	"time"
)

func TestTokenBucketAllow(t *testing.T) {
	tests := []struct {
		name           string
		capacity       int
		refillInterval time.Duration
		allowCalls     int
		sleepBetween   time.Duration
		expected       []bool
	}{
		{
			name:           "consume all tokens without refill",
			capacity:       5,
			refillInterval: time.Second,
			allowCalls:     6,
			sleepBetween:   0,
			expected:       []bool{true, true, true, true, true, false},
		},
		{
			name:           "tokens refill over time",
			capacity:       2,
			refillInterval: 2 * time.Second,
			allowCalls:     4,
			sleepBetween:   1 * time.Second,
			expected:       []bool{true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := NewTokenBucketRateLimiter(tt.capacity, tt.refillInterval)
			for i := 0; i < tt.allowCalls; i++ {
				time.Sleep(tt.sleepBetween)
				got := rc.Allow()
				if got != tt.expected[i] {
					t.Errorf("Call %d: expected %v, got %v", i+1, tt.expected[i], got)
				}
			}
		})
	}
}

func TestTokenBucketConcurrency(t *testing.T) {
	rc := NewTokenBucketRateLimiter(5, time.Second)
	results := make(chan bool, 20)
	for i := 0; i < 20; i++ {
		go func() {
			results <- rc.Allow()
		}()
	}
	allowed := 0
	for i := 0; i < 20; i++ {
		if <-results {
			allowed++
		}
	}
	if allowed > 5 {
		t.Errorf("More requests allowed than capacity! allowed=%d", allowed)
	}
}
