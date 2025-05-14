package advanced

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	tokens int
	leakRate time.Duration
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		tokens:   capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsed / lb.leakRate)
	lb.tokens += tokensToAdd
	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}
	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	fmt.Printf("Tokens Added: %d, Total Tokens: %d\n", tokensToAdd, lb.tokens)

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func main() {
	bucket := NewLeakyBucket(3, 200*time.Millisecond)
	var wg sync.WaitGroup
	var printMu sync.Mutex // Mutex for synchronized printing

	// Create 20 concurrent requests
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			
			// Random sleep between 0-300ms to simulate random request timing
			time.Sleep(time.Duration(requestID) * time.Millisecond * 50)
			
			if bucket.Allow() {
				printMu.Lock()
				fmt.Printf("Request %d: Allowed\n", requestID)
				printMu.Unlock()
			} else {
				printMu.Lock()
				fmt.Printf("Request %d: Blocked\n", requestID)
				printMu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\nAll requests completed")
}
