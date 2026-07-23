package waitgroup

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize(){
	fmt.Println("This will not be repeated more than once")

}

func SyncOnceDemo(){
	var wg sync.WaitGroup
	for i := range 5{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("Go routine #",i)
			once.Do(initialize)
		}()
	}
	wg.Wait()
}