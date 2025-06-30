package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	fmt.Printf("Worker %d task completed\n", id)
}


func main() {
	fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup

	for i := 1; i < 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("workers task done")		
}