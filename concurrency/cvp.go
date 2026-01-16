package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting\n", id)
	for range 100_000_000 {

	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d is finished \n", id)
}

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup
	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}
	wg.Wait()
}
