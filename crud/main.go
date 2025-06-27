package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("We are learning HTTP requests in Go")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", res.StatusCode)
		return
	}
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Response Data:", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo:", todo)
	fmt.Println("User ID:", todo.UserID)
	fmt.Println("ID:", todo.Id)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Todo:", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserID:    1,
		Id:        2,
		Title:     "Hi Mom!",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error: received non-201 status code:", res.StatusCode)
		return
	}
	var createdTodo Todo
	err = json.NewDecoder(res.Body).Decode(&createdTodo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Created Todo:", createdTodo)
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    1,
		Id:        2,
		Title:     "Updated Title",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	defer req.Body.Close()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", res.StatusCode)
		return
	}
	var updatedTodo Todo
	err = json.NewDecoder(res.Body).Decode(&updatedTodo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Updated Todo:", updatedTodo)
}


func performDeleteRequest() {

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/1", nil)

	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making DELETE request:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", res.StatusCode)
		return
	}
	fmt.Println("Todo deleted successfully")
	fmt.Println("Response Status:", res.Status)
}