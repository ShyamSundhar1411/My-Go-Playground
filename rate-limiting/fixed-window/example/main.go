package main

import (
	"fmt"
	"time"

	fixedwindow "github.com/ShyamSundhar1411/My-Go-Playground/rate-limiting/fixed-window"
)

func main() {
	rc := fixedwindow.NewFixedWindowRateLimiter(5, time.Second)
	fmt.Println("Fixed Window Rate Limiter")
	fmt.Println(rc)
	for i := range 10 {
		if rc.Allow() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "blocked")
		}
	}
}
