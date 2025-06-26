package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")

	myUrl := "https://www.google.com/search?q=golang+url+parsing&hl=en"

	parsedUrl, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Query:", parsedUrl.RawQuery) // RawQuery is the encoded query string
	fmt.Println("URL:", parsedUrl)
}
