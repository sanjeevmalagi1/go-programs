package main

import (
	"fmt"
	"io"
	"net/http"
)

var base_url = "https://jsonplaceholder.typicode.com"

func main() {
	// Define the API endpoint
	url := base_url + "/posts"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
