package conditions

import (
	"fmt"
	"sync"
	"time"
)

func HoldAndWait() {
	var resourceA, resourceB sync.Mutex

	go func() {
		resourceA.Lock()
		fmt.Println("Goroutine 1: locking Resource A")
		time.Sleep(time.Second * 1)
		fmt.Println("Goroutine 1: trying to accquire Resource B")
		resourceB.Lock()

		fmt.Println("Goroutine 1: acquired both resources")
		resourceA.Unlock()
		resourceB.Unlock()
		fmt.Println("Goroutine 1 finished")
	}()

	go func() {
		resourceB.Lock()
		defer resourceB.Unlock()
		fmt.Println("Goroutine 2: locking Resource B")
		fmt.Println("Goroutine 2: holding Resource B indefinitely")
		select {}

	}()
	select {}
}
