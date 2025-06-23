package main

import "fmt"

func main() {

	// Define a struct to represent a person
	type Person struct {
		Name    string
		Age     int
		isAlive bool
	}

	// // Create an instance of the Person struct
	// alice := Person{Name: "Alice", Age: 30, isAlive: true}

	// // Alternatively, you can create an instance without using a composite literal
	// var prajualit Person
	// // or
	// // prajualit := Person{}
	// prajualit.Name = "Prajualit"
	// prajualit.Age = 25
	// prajualit.isAlive = true

	// // new keyword
	// arav := new(Person)
	// arav.Name = "Arav"
	// arav.Age = 28
	// arav.isAlive = true

	// // Print the person's details
	// fmt.Println("Person Details:", alice)
	// fmt.Printf("Name: %s, Age: %d, isAlive: %t", alice.Name, alice.Age, alice.isAlive)
	// fmt.Printf("\nName: %s, Age: %d, isAlive: %t", prajualit.Name, prajualit.Age, prajualit.isAlive)
	// fmt.Printf("\nName: %s, Age: %d, isAlive: %t", arav.Name, arav.Age, arav.isAlive)

	// we can also use a struct field in another struct
	type Address struct {
		Street string
		City   string
		State  string
		Zip    string
	}
	type Employee struct {
		Person
		EmployeeID string
		Department string
		Address    // Embedded struct
	}

	// Create an instance of the Employee struct
	john := Employee{
		Person:     Person{Name: "John", Age: 35, isAlive: true},
		EmployeeID: "E12345",
		Department: "Engineering",
		Address: Address{
			Street: "123 Main St",
			City:   "Springfield",
			State:  "IL",
			Zip:    "62701",
		},
	}
	// Print the employee's details
	fmt.Println("Employee Details:")
	fmt.Println("Employee :", john)
	fmt.Printf("Name: %s, Age: %d, isAlive: %t\n", john.Name, john.Age, john.isAlive)
	fmt.Printf("Employee ID: %s, Department: %s\n", john.EmployeeID, john.Department)
	fmt.Printf("Address: %s, %s, %s, %s\n", john.Address.Street, john.Address.City, john.Address.State, john.Address.Zip) 
	fmt.Printf("Address: %s, %s, %s, %s\n", john.Street, john.City, john.State, john.Zip)// Accessing embedded struct fields directly
}
