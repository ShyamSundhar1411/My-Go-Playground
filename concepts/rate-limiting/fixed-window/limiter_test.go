package fixedwindow

import (
	"sync"
	"testing"
	"time"
)

func TestFixedWindowAllo(t *testing.T) {
	tests := []struct {
		name         string
		limit        int
		windowSize   time.Duration
		allowedCalls int
		sleepBetween time.Duration
		expected     []bool
	}{
		{
			name:         "allow within same window and reject after limit",
			limit:        3,
			windowSize:   2 * time.Second,
			allowedCalls: 5,
			sleepBetween: 0,
			expected:     []bool{true, true, true, false, false},
		},
		{
			name:         "window resets after expiry",
			limit:        2,
			windowSize:   1 * time.Second,
			allowedCalls: 4,
			sleepBetween: 500 * time.Millisecond,
			expected:     []bool{true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewFixedWindowRateLimiter(tt.limit, tt.windowSize)
			for i := 0; i < tt.allowedCalls; i++ {
				if i > 0 {
					time.Sleep(tt.sleepBetween)
				}
				allowed := limiter.Allow()
				if allowed != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected[i], allowed)
				}
			}
		})
	}
}
func TestFixedWindowConcurrency(t *testing.T) {
	limit := 10
	windowSize := 5 * time.Second
	rc := NewFixedWindowRateLimiter(limit, windowSize)
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
