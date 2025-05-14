package advanced

import (
	"fmt"
	"time"
)

//

type tickerRequest struct {
	personId int
	numTicks int
	cost     int
}

func tickerWorker(tickerRequests <-chan tickerRequest, results chan<- int) {
	for req := range tickerRequests {
		fmt.Printf("Processing request for person %d with %d ticks and cost %d\n", req.personId, req.numTicks, req.cost)
		time.Sleep(time.Second * 1)
		results <- req.cost
	}
}

func main() {
	numRequests := 5
	pricePerTick := 10

	requests := make(chan tickerRequest, numRequests)
	results := make(chan int, numRequests)

	for i := 0; i < 3; i++ {
		go tickerWorker(requests, results)
	}

	for i := 0; i < numRequests; i++ {
		requests <- tickerRequest{
			personId: i + 1,
			numTicks: (i + 1) * 2,
			cost:     pricePerTick * (i + 1) * 2,
		}
	}
	close(requests)

	for i := 0; i < numRequests; i++ {
		result := <-results
		fmt.Printf("The ticket request for person %d has been processed with cost %d\n", i+1, result)
	}
}



