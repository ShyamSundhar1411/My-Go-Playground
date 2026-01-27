package main

import (
	"fmt"
	"os"

	"github.com/ShyamSundhar1411/My-Go-Playground/concurrency/deadlocks/conditions"
)

func main() {
	fmt.Println("Deadlock Condition Simulator")
	fmt.Println("----------------------------")
	fmt.Println("1. Circular Wait")
	fmt.Println("2. Hold and Wait")
	fmt.Println("3. Mutual Exclusion")
	fmt.Println("4. No Preemption")
	fmt.Print("\nSelect an option: ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	switch choice {
	case 1:
		fmt.Println("Circular Wait Condition")
		conditions.CircularWait()
	case 2:
		fmt.Println("Hold and Wait Condition")
		conditions.HoldAndWait()
	case 3:
		fmt.Println("Mutual Exclusion Condition")
		conditions.MutualExclusion()
	case 4:
		fmt.Println("No Preemption Condition")
		conditions.NoPreemption()
	default:
		fmt.Println("Invalid Option")
	}

}
