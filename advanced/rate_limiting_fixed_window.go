package advanced

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	count      int
	limit      int
	window time.Duration
	resetTime  time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit: limit,
		window: window,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.After(r.resetTime) {
		r.count = 0
		r.resetTime = now.Add(r.window)
	}
	if r.count < r.limit {
		r.count++
		return true
	}
	return false
}


func main() {

	limiter := NewRateLimiter(3, 3*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			time.Sleep(time.Duration(index) * 500 * time.Millisecond)
			if limiter.Allow() {
				fmt.Printf("Request %d allowed\n", index)
			} else {
				fmt.Printf("Request %d blocked\n", index)
			}
		}(i)
	}
	wg.Wait()  // Wait for all goroutines to complete

	


}
