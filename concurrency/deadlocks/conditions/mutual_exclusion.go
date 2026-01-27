package conditions

import (
	"fmt"
	"sync"
)

func MutualExclusion() {
	var resourceA sync.Mutex

	go func() {
		resourceA.Lock()
		defer resourceA.Unlock()
		fmt.Println("Goroutine 1 locked Resource A")
		fmt.Println("Goroutine 1 using Resource A indefinitely")
		select {}

	}()

	go func() {
		fmt.Println("Goroutine 2: trying to lock Resource A")
		resourceA.Lock()
		defer resourceA.Unlock()
		fmt.Println("Goroutine 2 locked Resource A")
	}()
	select {}
}
