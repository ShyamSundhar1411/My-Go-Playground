package main

import (
	"fmt"
	"time"

	leakybucket "github.com/ShyamSundhar1411/My-Go-Playground/rate-limiting/leaky-bucket"
)

func main() {
	rc := leakybucket.NewLeakyBucketRateLimiter(5, time.Second)
	fmt.Println("Leaky Bucket Rate Limiter")
	fmt.Println(rc)
	for i := range 10 {
		if rc.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "blocked")
		}
	}
}
