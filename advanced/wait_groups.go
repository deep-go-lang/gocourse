package advanced

import (
	"fmt"
	"sync"
	"time"
)

type tickerRequest struct {
	personId int
	numTicks int
	cost     int
}

func tickerWorker(wg *sync.WaitGroup, tickerRequests <-chan tickerRequest, results chan<- int) {
	defer wg.Done()
	for req := range tickerRequests {
		fmt.Printf("Processing request for person %d with %d ticks and cost %d\n", req.personId, req.numTicks, req.cost)
		time.Sleep(time.Second * 1)
		results <- req.cost
	}
}

func main() {
	numRequests := 5
	pricePerTick := 10
	numWorkers := 3

	requests := make(chan tickerRequest, numRequests)
	results := make(chan int, numRequests)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go tickerWorker(&wg, requests, results)
	}

	// Send requests
	for i := 0; i < numRequests; i++ {
		requests <- tickerRequest{
			personId: i + 1,
			numTicks: (i + 1) * 2,
			cost:     pricePerTick * (i + 1) * 2,
		}
	}
	close(requests)

	// Wait for all workers to finish in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("The ticket request has been processed with cost %d\n", result)
	}
}

// func main() {

// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := 0; i < numWorkers; i++ {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("All workers finished")

// }

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)

// 	// Process jobs until the jobs channel is closed
// 	for job := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, job)
// 		time.Sleep(time.Second * 2) // Simulate work
// 		results <- job * 2
// 		fmt.Printf("Worker %d finished job %d\n", id, job)
// 	}

// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5

// 	// Create channels for jobs and results
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// Start workers
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go worker(i+1, &wg, jobs, results)
// 	}

// 	// Send jobs to workers
// 	for i := 0; i < numJobs; i++ {
// 		jobs <- i+1
// 	}
// 	close(jobs) // Close jobs channel to signal no more jobs

// 	// Wait for all workers to finish in a separate goroutine
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Collect results
// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// }

// type Worker struct {
// 	ID int
// 	Task string
// }

// func (w *Worker) Work(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting task: %s\n", w.ID, w.Task)
// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("Worker %d finished task: %s\n", w.ID, w.Task)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	tasks := []string{"digging", "cleaning", "cooking", "washing", "shopping"}

// 	for i, task := range tasks {
// 		worker := &Worker{ID: i + 1, Task: task}
// 		wg.Add(1)
// 		go worker.Work(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers finished")

// }

