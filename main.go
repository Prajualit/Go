package main

import (
	"fmt"
	"myLearning/myUtil"
)

func main() {
    fmt.Println("Hi Mom!")
    fmt.Println("Welcome to myLearning project!")
	myUtil.PrintMessage("Hello from myUtil package!")

	var name string = "Prajualit"
	fmt.Println("My name is", name)

	slogan := "Hi Mom! I am learning Go!"
	fmt.Println(slogan)
	
	Public := "This is a public variable"
	private := "This is a private variable"
	// Public variable can be accessed outside the package
	// private variable cannot be accessed outside the package
	// Same is the case with functions, types, and methods
	fmt.Println("Public variable:", Public)
	fmt.Println("Private variable:", private)

	var name1 string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name1)
	fmt.Println("Hello", name1, "! Welcome to myLearning project!")
	// the scan function reads only until the first space, so if you enter a name with a space, it will only read the first part
	// we can use fmt.Scanf to read the whole line
}