package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 10
	fmt.Printf("Original number type: %T", num)

	num_str := strconv.Itoa(num) // I(integer)  to  A(abcd)
	fmt.Printf("\nConverted number to string: %s, type: %T", num_str, num_str)

	num1, _ := strconv.Atoi(num_str) // A(abcd) to I(integer)
	fmt.Printf("\nConverted string back to number: %d, type: %T", num1, num1)

	// float
	num_float, _ := strconv.ParseFloat(num_str, 64) 
	fmt.Printf("\nConverted string to float: %f, type: %T", num_float, num_float)
}
