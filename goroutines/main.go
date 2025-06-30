package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, Goroutine!")
	time.Sleep(2 * time.Second) // Sleep for 2 seconds to simulate work
	fmt.Println("Goroutine 1 finished execution")
}

func sayhi() {
	fmt.Println("Hi, Goroutine!")
	time.Sleep(1 * time.Second) // Sleep for 1 second to simulate work
	fmt.Println("Goroutine 2 finished execution")
}

func main() {
	fmt.Println("We are learning Goroutines in Go")
	go sayHello()
	go sayhi()
	fmt.Println("Main function is running")

	time.Sleep(1 * time.Second) // Wait for goroutine to finish
	fmt.Println("Main function finished execution")
}
