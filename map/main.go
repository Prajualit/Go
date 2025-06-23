package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 92
	studentGrades["Diana"] = 88
	studentGrades["Eve"] = 95
	studentGrades["Frank"] = 80

	fmt.Println("Bob Grades:", studentGrades["Bob"])
	
	delete(studentGrades, "Alice")

	// The existence of a key is checked using the second return value of a map lookup.
	grades, exists := studentGrades["Alice"]
	fmt.Println("Alice Grades:", grades, "Exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("Index: %s, Value: %d\n", index, value)
		// data is stored in an unordered manner due to which the order of iteration is mixed up everytime you run the program.
	}
	
	person := map[string]int{
		"John": 78,
		"Jane": 85,
		"Mike": 90,
		"Lucy": 88,
	}
	
	fmt.Println("Person Grades:", person)
	for index, value := range person {
		fmt.Printf("Index: %s, Value: %d\n", index, value)
	}
}
