package main

import (
	"fmt"
	"os"
)

func main() {

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }

	// fmt.Println(file.Name())
	// defer file.Close()
	// // file needs to be closed since it can cause a resource leak if not closed properly
	// fmt.Println("File created successfully, and it will be closed at the end of the program.")

	// content := "Hello, World!\nThis is a sample text file.\n"
	// byte, errs := io.WriteString(file, content)
	// fmt.Printf("Wrote %d bytes to the file.\n", byte)
	// if errs != nil {
	// 	fmt.Println("Error writing to file:", errs)
	// 	return
	// }

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)

	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break // End of file reached
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error reading file:", err)
	// 		return
	// 	}

	// 	fmt.Println("File content read successfully:")
	// 	fmt.Println(string(buffer[:n]))
	// }

	content, err := os.ReadFile("example.txt")
	// this function is not used if we have a big file since it reads the whole file into memory
	// if the file is too big, it can cause memory issues
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content read successfully:")
	fmt.Println(string(content))
}
