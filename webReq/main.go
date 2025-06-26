package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Learning web services in Go!")
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	fmt.Println("Response Status:", res.Status)
	defer res.Body.Close() // Ensure the response body is closed after we're done with it

	data, err := io.ReadAll(res.Body) // Read the response body to completion
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body Length:", len(data))
	fmt.Println("Response Body Content:", string(data[:100])) // Print the first 100 characters
}
