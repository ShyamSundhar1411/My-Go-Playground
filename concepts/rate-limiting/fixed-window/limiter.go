package fixedwindow

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	limit      int
	count      int
	windowSize time.Duration
	startTime  time.Time
	mu         sync.Mutex
}

func NewFixedWindowRateLimiter(limit int, windowSize time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		limit:      limit,
		count:      0,
		windowSize: windowSize,
		startTime:  time.Now(),
	}
}

func (rc *FixedWindowRateLimiter) Allow() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	currentTime := time.Now()
	if currentTime.Sub(rc.startTime).Seconds() >= rc.windowSize.Seconds() {
		rc.count = 0
		rc.startTime = currentTime
		fmt.Printf(
			"[Reset] Window expired | Previous count=%d | New window start=%v\n",
			rc.count,
			currentTime.Format("15:04:05.000"),
		)
	}
	if rc.count < rc.limit {
		rc.count++
		fmt.Printf(
			"[Allow] Time=%s | Count=%d/%d\n",
			currentTime.Format("15:04:05.000"),
			rc.count,
			rc.limit,
		)
		return true
	}
	fmt.Printf(
		"[Reject] Time=%s | Count=%d/%d\n",
		currentTime.Format("15:04:05.000"),
		rc.count,
		rc.limit,
	)
	return false
}

func (rc *FixedWindowRateLimiter) String() string {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	return fmt.Sprintf(
		"FixedWindow{limit=%d, count=%d, windowSize=%v, startTime=%v}",
		rc.limit,
		rc.count,
		rc.windowSize,
		rc.startTime,
	)
}
