package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return a / b, fmt.Errorf("cannot divide %f by zero", b)
	}
	return a / b, nil
}

func main() {
	fmt.Println("error handling in Go")
	// ans, _ := divide(10, 4) //this will ignore the error
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("The answer is:", ans)
}
