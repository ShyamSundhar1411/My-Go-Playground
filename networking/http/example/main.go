package main

import (
	"fmt"
	"os"

	"github.com/ShyamSundhar1411/My-Go-Playground/networking/http"
)

func main() {
	fmt.Println("HTTP Simulator")
	fmt.Println("----------------------------")
	fmt.Println("1. Standalone HTTP Server")
	fmt.Println("2. Standalone HTTP Client")
	fmt.Println("3. Server + Client Demo")
	fmt.Print("\nSelect an option: ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	switch choice {
	case 1:
		fmt.Println("Starting Standalone HTTP Server")
		http.HttpServer()
	case 2:
		fmt.Println("Starting Standalone HTTP Client")
		http.HttpClient()
	default:
		fmt.Println("Invalid Option")
	}
}
