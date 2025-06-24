package main

import "fmt"

func modifyValueByReference(x *int) {
	*x = *x + 20
}

func main() {

	num := 10
	ptr := &num
	// var ptr *int // Allocating memory for an int and assigning it to ptr
	// ptr = &num // Assigning the address of num to ptr
	
	// another way to create a pointer is to use the new keyword
	// ptr = new(int)

	fmt.Println("num:", num)
	fmt.Println("ptr:", ptr)
	fmt.Println("ptr:", *ptr) // Dereferencing the pointer to get the value it points to

	var ptr1 *int

	fmt.Println("ptr1:", ptr1) // This will print <nil> since ptr1 is not initialized

	modifyValueByReference(&num)
	fmt.Println("After modification, num:", num)
	
}
