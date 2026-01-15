package slidingwindowlog

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLogRateLimiter struct {
	limit       int
	windowSize  time.Duration
	requestsLog []time.Time
	mu          sync.Mutex
}

func NewSlidingWindowLogRateLimiter(limit int, windowSize time.Duration) *SlidingWindowLogRateLimiter {
	return &SlidingWindowLogRateLimiter{
		limit:       limit,
		windowSize:  windowSize,
		requestsLog: make([]time.Time, 0),
	}
}

func (rc *SlidingWindowLogRateLimiter) Allow() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	currentTime := time.Now()
	for len(rc.requestsLog) > 0 && currentTime.Sub(rc.requestsLog[0]) >= rc.windowSize {
		fmt.Printf(
			"[Updating Log] Removing Logs | Previous count=%d | Log=%v",
			len(rc.requestsLog),
			rc.requestsLog,
		)
		rc.requestsLog = rc.requestsLog[1:]
	}
	if len(rc.requestsLog) < rc.limit {
		rc.requestsLog = append(rc.requestsLog, currentTime)
		fmt.Printf(
			"[Allow] Adding Logs | Previous count=%d | Log=%v",
			len(rc.requestsLog),
			rc.requestsLog,
		)
		return true
	}
	fmt.Printf(
		"[Reject]  Previous count=%d | Log=%v",
		len(rc.requestsLog),
		rc.requestsLog,
	)
	return false
}

func (rc *SlidingWindowLogRateLimiter) String() string {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return fmt.Sprintf(
		"SlidingWindowLog{limit=%d, windowSize=%s, requestsLog=%v}",
		rc.limit,
		rc.windowSize,
		rc.requestsLog,
	)
}
