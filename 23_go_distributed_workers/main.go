package main

import (
	"fmt"
	"sync"
	"time"
)

// Job struct defines the structure of a job.
type Job struct {
	ID     int
	Data   string
	Result string
}

// Worker function processes jobs and sends results back to the results channel.
func worker(id int, jobs <-chan Job, results chan<- Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate job processing (e.g., reversing the Data string).
		fmt.Printf("Worker %d processing job ID %d\n", id, job.ID)
		job.Result = reverseString(job.Data)
		results <- job
	}
}

// reverseString is a helper function to simulate job processing.
func reverseString(s string) string {
	time.Sleep(1 * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	const numJobs = 10
	const numWorkers = 5

	jobsCh := make(chan Job, numJobs)    // Channel to send jobs to workers.
	resultsCh := make(chan Job, numJobs) // Channel to collect results from workers.
	var wg sync.WaitGroup                // WaitGroup to ensure all workers finish.

	// Start worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobsCh, resultsCh, &wg)
	}

	// Send jobs to the jobs channel.
	for i := 1; i <= numJobs; i++ {
		jobsCh <- Job{ID: i, Data: fmt.Sprintf("Data%d", i)}
	}
	close(jobsCh) // Close the jobs channel to indicate no more jobs.

	// Wait for all workers to finish.
	go func() {
		wg.Wait()
		close(resultsCh) // Close the results channel once all workers are done.
	}()

	fmt.Println("After wg")

	// Collect results and store them in a map.
	jobResults := make(map[int]string)
	for result := range resultsCh {
		jobResults[result.ID] = result.Result
	}

	fmt.Println("\nJob Results:")
	for id, result := range jobResults {
		fmt.Printf("Job ID %d: Result = %s\n", id, result)
	}
}
