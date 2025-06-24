package main

import "fmt"

func main() {
	fmt.Println("Starting the program...")
	defer fmt.Println("This will always execute at the end, even if there's an error or panic.")
	fmt.Println("Performing some operations...")
	// Simulating an operation that might cause a panic
	// Uncomment the next line to see how defer works with a panic
	// panic("Something went wrong!")
	fmt.Println("Operations completed successfully.")
}
