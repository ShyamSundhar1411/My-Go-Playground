package waitgroup

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func SyncPoolDemo() {
	var pool = sync.Pool{
		New: func() interface{}{
			fmt.Println("Creating a new person")
			return &person{}
		},

	}
	p1 := pool.Get().(*person)
	p1.name = "Person 1"
	p1.age = 18
	fmt.Println("Got person",p1)
	pool.Put(p1)
	fmt.Println("Put person back")

	p2 := pool.Get().(*person)
	p2.name = "Person 2"
	p2.age = 21
	fmt.Println("Got person", p2)
	pool.Put(p2)
	
}