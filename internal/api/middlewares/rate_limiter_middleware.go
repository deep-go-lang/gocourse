package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu sync.Mutex
	visitors map[string]int
	limit int
	resetInterval time.Duration
}

func NewRateLimiter(limit int, resetInterval time.Duration) *RateLimiter {
	m := &RateLimiter{
		limit: limit,
		resetInterval: resetInterval,
		visitors: make(map[string]int),
	}
	go m.resetVisitors()
	return m
}

func (m *RateLimiter) resetVisitors() {
	for {
		time.Sleep(m.resetInterval)
		m.mu.Lock()
		m.visitors = make(map[string]int)
		m.mu.Unlock()
	}
}

func (m *RateLimiter) Middleware(next http.Handler) http.Handler {
	fmt.Println("RateLimiter Middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RateLimiter Middleware being returned...")
		m.mu.Lock()
		defer m.mu.Unlock()

		ip := r.RemoteAddr
		currentCount := m.visitors[ip]
		
		if currentCount >= m.limit {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		
		m.visitors[ip]++
		fmt.Printf("Visitor count from %s: %d\n", ip, m.visitors[ip])
		next.ServeHTTP(w, r)
		fmt.Println("RateLimiter Middleware serving next handler...")
	})
}
