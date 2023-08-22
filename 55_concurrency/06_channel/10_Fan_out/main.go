package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5) // Buffered channel for jobs
	results := make(chan int) // Channel for results

	// Number of worker goroutines
	numWorkers := 3

	// Launch worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to workers
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate some processing time
		result := job * 2
		fmt.Printf("Worker %d: Job %d processed\n", id, job)
		results <- result
	}
}
