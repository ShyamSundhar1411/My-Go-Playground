package conditions

import (
	"fmt"
	"sync"
	"time"
)

func NoPreemption() {
	var resourceA sync.Mutex

	go func(){
		resourceA.Lock()
		defer resourceA.Unlock()
		fmt.Println("Goroutine 1: accquired Resource A")
		time.Sleep(10*time.Second)
		fmt.Println("Goroutine 1: Releasing Resource A")
	}()
	time.Sleep(1*time.Second)
	go func(){
		fmt.Println("Goroutine 2: Trying to accquire Resource A")
		resourceA.Lock()
		defer resourceA.Unlock()
		fmt.Println("Goroutine 2: accquired Resource A")
	}()
	select {}
}