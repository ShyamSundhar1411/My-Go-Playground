package main

import (
	"fmt"
	"time"

	slidingwindowlog "github.com/ShyamSundhar1411/My-Go-Playground/rate-limiting/sliding-window-log"
)

func main() {
	rc := slidingwindowlog.NewSlidingWindowLogRateLimiter(3, time.Second)
	fmt.Println("Sliding Window Log Rate Limiter")
	fmt.Println(rc)
	for i := range 10 {
		if rc.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "blocked")
		}
	}
}
