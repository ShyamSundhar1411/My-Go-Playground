package leakybucket

import (
	"testing"
	"time"
)

func TestLeakyBucketAllow(t *testing.T) {
	tests := []struct {
		name         string
		capacity     int
		leakInterval time.Duration
		allowedCalls int
		sleepBetween time.Duration
		expected     []bool
	}{
		{
			name:         "bucket fills up and rejects excess",
			capacity:     3,
			leakInterval: 10 * time.Second,
			allowedCalls: 5,
			sleepBetween: 0,
			expected:     []bool{true, true, true, false, false},
		},
		{
			name:         "bucket fills up and leaks",
			capacity:     2,
			leakInterval: 2 * time.Second,
			allowedCalls: 4,
			sleepBetween: 1 * time.Second,
			expected:     []bool{true, true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewLeakyBucketRateLimiter(tt.capacity, tt.leakInterval)
			for i, expected := range tt.expected {
				actual := limiter.Allow()
				if actual != expected {
					t.Errorf("call %d: expected %t, got %t", i, expected, actual)
				}
				if i < len(tt.expected)-1 {
					time.Sleep(tt.sleepBetween)
				}
			}
		})

	}
}

func TestLeakyBucketConcurrency(t *testing.T) {
	rc := NewLeakyBucketRateLimiter(5, time.Second)
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
	if allowed != 5 {
		t.Errorf("expected 5 allowed calls, got %d", allowed)
	}
}
