package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	counter int64
	mu      sync.Mutex
}

func (ac *AtomicCounter) Increment() {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	atomic.AddInt64(&ac.counter, 1)
}

func (ac *AtomicCounter) Decrement() {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	atomic.AddInt64(&ac.counter, -1)
}

func (ac *AtomicCounter) Get() int64 {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	return atomic.LoadInt64(&ac.counter)
}

func main() {
	ac := &AtomicCounter{counter: 25}
	

	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				ac.Decrement()
				// value++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", ac.Get())
}
