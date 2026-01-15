package main

import (
	"fmt"
	"github.com/ShyamSundhar1411/My-Go-Playground/rate-limiting/token-bucket"
	"time"
)

func main() {
	rc := tokenbucket.NewTokenBucketRateLimiter(5, time.Minute)
	fmt.Println("Token Bucket Rate Limiter")
	fmt.Println(rc)
	for i := range 10 {
		if rc.Allow() {
			fmt.Printf("Request %d allowed\n", i+1)
		} else {
			fmt.Printf("Request %d rejected\n", i+1)
		}
	}
}
