package waitgroup

import (
	"fmt"
	"sync"
	"time"
)
 
const bufferSize = 5
type buffer struct{
	items [] int;
	mu sync.Mutex;
	cond *sync.Cond;
}

func newBuffer(size int) *buffer{
	b := &buffer{items: make([]int, 0,size)}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int){
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == bufferSize{
		b.cond.Wait()	
	}
	b.items = append(b.items, item)
	fmt.Println("Produced item:",item)
	b.cond.Signal()
}

func (b *buffer) consume(){
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0{
		b.cond.Wait()
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed item:", item)
	b.cond.Signal()
}

func producer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()
	for i := range(10){
		b.produce(i+100)
		time.Sleep(time.Millisecond*100)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()
	for range(10){
		b.consume()
		time.Sleep(time.Millisecond*300)
	}
}

func ProducerConsumerDemo(){
	var wg sync.WaitGroup
	b := newBuffer(bufferSize)
	wg.Add(2)
	go producer(b, &wg)
	go consumer(b, &wg)
	wg.Wait()
}