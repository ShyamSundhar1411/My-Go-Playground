package main

import (
	"fmt"
	"os"

	"github.com/ShyamSundhar1411/My-Go-Playground/concurrency/synchronization/waitgroup"
)

func main() {
	fmt.Println("Wait Group Simulator")
	fmt.Println("----------------------------")
	fmt.Println("1. Sync Conditions")
	fmt.Println("2. Sync Once")
	fmt.Println("3. Sync Pool")
	fmt.Print("\nSelect an option: ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	switch choice{
	case 1:
		fmt.Println("Sync Conditions")
		waitgroup.ProducerConsumerDemo()
	case 2:
		fmt.Println("Sync Once")
		waitgroup.SyncOnceDemo()
	case 3:
		fmt.Println("Sync Pool")
		waitgroup.SyncPoolDemo()
	default:
		fmt.Println("Invalid Option")
	}
}