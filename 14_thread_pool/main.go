package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate time-consuming work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Process result and send it to results channel
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)    // Job queue channel
	results := make(chan int, numJobs) // Result channel

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
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}
