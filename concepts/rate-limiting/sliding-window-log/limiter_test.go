package slidingwindowlog

import (
	"sync"
	"testing"
	"time"
)

func TestSlidingWindowLogAllow(t *testing.T) {
	tests := []struct {
		name         string
		limit        int
		windowSize   time.Duration
		allowedCalls int
		expected     []bool
		sleepBetween time.Duration
	}{
		{
			name:         "window fills up and denies the requests",
			limit:        3,
			windowSize:   time.Second,
			allowedCalls: 5,
			expected:     []bool{true, true, true, false, false},
		},
		{
			name:         "window cleans logs and accepts every requests",
			limit:        3,
			windowSize:   time.Second,
			allowedCalls: 5,
			expected:     []bool{true, true, true, true, true},
			sleepBetween: time.Second * 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewSlidingWindowLogRateLimiter(tt.limit, tt.windowSize)
			for i := 0; i < tt.allowedCalls; i++ {
				if tt.sleepBetween != 0 {
					time.Sleep(tt.sleepBetween)
				}
				allowed := limiter.Allow()
				if allowed != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected[i], allowed)
				}
			}
		})

	}
}

func TestSlidingWindowLogConcurrency(t *testing.T) {
	limit := 10
	windowSize := time.Second * 5
	rc := NewSlidingWindowLogRateLimiter(limit, windowSize)
	var wg sync.WaitGroup
	totalGoroutines := 100
	results := make(chan bool, totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- rc.Allow()
		}()
	}
	wg.Wait()
	close(results)
	allowed := 0
	for r := range results {
		if r {
			allowed++
		}
	}
	if allowed > limit {
		t.Errorf(
			"concurrency violation: allowed %d requests, limit is %d",
			allowed,
			limit,
		)
	}
}
