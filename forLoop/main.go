package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	counter := 0
	for {
		if counter >= 10 {
			break
		}
		fmt.Println("Counter:", counter)
		counter++
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	text := "Hi Mom!"
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
