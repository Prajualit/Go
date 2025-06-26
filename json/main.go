package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("We are learning JSON")
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St, Anytown, USA",
	}
	fmt.Println("Person:", person)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Decoded Person:", decodedPerson)
}
