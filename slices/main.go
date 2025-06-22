package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 9, 10)
	
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length of numbers slice:", len(numbers))
	fmt.Println("Capacity of numbers slice:", cap(numbers))
	// difference between length and capacity:
	// Length is the number of elements in the slice.
	// Capacity is the number of elements the slice can hold before it needs to be resized.
	
	
	numbers1 := make([]int, 5, 10)
	// The above line creates a slice of integers with an initial length of 5 and a capacity of 10.
	
	numbers1 = append(numbers1, 11, 12)
	// The above line appends 11 and 12 to the slice, increasing its length.
	numbers1[2] = 6
	// The above line sets the third element of the slice to 6.

	fmt.Println("Numbers1:", numbers1)
	fmt.Println("Length of numbers slice:", len(numbers1))
	fmt.Println("Capacity of numbers slice:", cap(numbers1))

}