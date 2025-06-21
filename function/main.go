package main

import (
	"fmt"
)

func firstFunction() {
	fmt.Println("This is the first function in main package.")
}

func add(a int, b int) (int) {
	return a + b
}

func main() {
	fmt.Println("Hi Mom!")
	firstFunction()
	fmt.Println(add(5, 3)) // This will call the add function and print the result
}