package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, banana, cherry, date"
	dataSlice := strings.Split(data, ", ")
	fmt.Println("Split string into slice:", dataSlice) // [apple banana cherry date]

	str := "1 2 3 45 3234134123"
	strSlice := strings.Count(str, "3")
	fmt.Println("Count of spaces in string:", strSlice) // 5

	joinedStr := strings.Join(dataSlice, " | ")
	fmt.Println("Joined string:", joinedStr) // apple | banana | cherry | date
}
