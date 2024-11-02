package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var base_url = "https://jsonplaceholder.typicode.com"

func make_api_request(count int) (string, error) {
	url := base_url + "/posts"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return "", err
	}

	fmt.Println("API request successful :", count)
	return string(body), nil
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(count int) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			make_api_request(count)
		}(i)
	}

	wg.Wait()
	// fmt.Println("Response:", string(body))
}
