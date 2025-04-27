package advanced

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	requests chan struct{}
	interval time.Duration
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	r := &RateLimiter{requests: make(chan struct{}, limit), interval: interval}
	for i := 0; i < limit; i++ {
		r.requests <- struct{}{}
	}
	go r.startRefill()
	return r
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.requests:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) startRefill() {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case r.requests <- struct{}{}:
			default:
				// Channel is full, drop the oldest request
			}
		}
	}	
}

func main() {

	rl := NewRateLimiter(5, time.Second)

	for i := 0; i < 10; i++ {
		if rl.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request blocked")
		}
		time.Sleep(300 * time.Millisecond)
	}

}
