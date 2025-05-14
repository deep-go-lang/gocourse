package advanced

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// StatefulWorker represents a worker that maintains state in a thread-safe manner
type StatefulWorker struct {
	mu      sync.Mutex
	count   int
	ch      chan int
	done    chan struct{}
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

// NewStatefulWorker creates a new StatefulWorker instance
func NewStatefulWorker() *StatefulWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &StatefulWorker{
		ch:     make(chan int, 100), // Buffered channel to prevent blocking
		done:   make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

// Start begins processing in a new goroutine
func (w *StatefulWorker) Start() {
	w.wg.Add(1)
	go w.process()
}

// process handles the main worker logic
func (w *StatefulWorker) process() {
	defer w.wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-w.ctx.Done():
			close(w.done)
			return
		case value := <-w.ch:
			w.mu.Lock()
			w.count += value
			fmt.Printf("Received value %d, new count: %d\n", value, w.count)
			w.mu.Unlock()
		case <-ticker.C:
			w.mu.Lock()
			w.count++
			fmt.Printf("Auto-increment, new count: %d\n", w.count)
			w.mu.Unlock()
		}
	}
}

// Send attempts to send a value to the worker
func (w *StatefulWorker) Send(value int) error {
	select {
	case w.ch <- value:
		return nil
	case <-w.ctx.Done():
		return fmt.Errorf("worker is shutting down")
	default:
		return fmt.Errorf("channel is full")
	}
}

// GetCount returns the current count value
func (w *StatefulWorker) GetCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.count
}

// Shutdown gracefully stops the worker
func (w *StatefulWorker) Shutdown() {
	w.cancel()
	w.wg.Wait() // Wait for all goroutines to complete
	close(w.ch)
}

func main() {
	// Create and start the worker
	w := NewStatefulWorker()
	w.Start()

	// Send some values
	for i := 0; i < 10; i++ {
		if err := w.Send(i); err != nil {
			fmt.Printf("Error sending value %d: %v\n", i, err)
			continue
		}
		time.Sleep(time.Second)
	}

	// Demonstrate error handling
	for i := 0; i < 1000; i++ {
		if err := w.Send(i); err != nil {
			fmt.Printf("Channel full, could not send value %d: %v\n", i, err)
			break
		}
	}

	// Get final count
	fmt.Printf("Final count: %d\n", w.GetCount())

	// Shutdown gracefully
	w.Shutdown()
	fmt.Println("Worker shutdown complete")
}
