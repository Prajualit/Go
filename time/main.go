package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time:", time.Now())
	formatted := time.Now().Format("January 2, 2006 15:04:05 MST") // "January 2, 2006 15:04:05 MST"
	// golange has defined date formats in and you cannot enter any other date
	// for example, if you want to format the date as "dd-mm-yyyy", you cannot do that directly
	// instead, you have to use the predefined layout
	// "02-01-2006" is the layout for "dd-mm-yyyy"
	// you canot use "02-01-2023" as it will not work
	// you have to use the predefined layout
	fmt.Println("Formatted time:", formatted)
}
