package main

import "fmt"

func main() {
	fmt.Println("hi Mom!")

	var name [5]string
	name[0] = "John"

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Names:", name)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length of names array:", len(name[0]))

	fmt.Printf("Name: %q\n", name[0])
	/* BONUS:
	if you want to print a quoted string
	*/
}
