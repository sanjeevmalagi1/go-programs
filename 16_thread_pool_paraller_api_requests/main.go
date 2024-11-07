package main

import (
	"fmt"
	"io"
	"net/http"
)

var base_url = "https://jsonplaceholder.typicode.com"

func make_api_request(count int) (string, error) {
	url := fmt.Sprintf("%s/posts/%d", base_url, count)

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

	return string(body), nil
}

func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		res, _ := make_api_request(job)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- res // Process result and send it to results channel
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 5

	jobs := make(chan int, numJobs)       // Job queue channel
	results := make(chan string, numJobs) // Result channel

	// Start a fixed number of workers (goroutines)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the job queue
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the job channel to indicate no more jobs

	// Collect results from the workers
	for a := 1; a <= numJobs; a++ {
		_ = <-results
		// fmt.Printf("Result: %d\n", result)
	}
}
