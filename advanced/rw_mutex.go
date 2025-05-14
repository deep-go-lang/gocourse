package advanced

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu   sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup, id int) {
	startTime := time.Now()
	rwmu.RLock()
	defer wg.Done()
	fmt.Printf("Reader %d: counter = %d (started at: %v)\n", id, counter, startTime.Format("15:04:05.000"))
	time.Sleep(100 * time.Millisecond) // Simulate some work
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	startTime := time.Now()
	rwmu.Lock()
	defer wg.Done()
	fmt.Printf("Writer: changing counter from %d to %d (started at: %v)\n", counter, value, startTime.Format("15:04:05.000"))
	counter = value
	time.Sleep(200 * time.Millisecond) // Simulate some work
	rwmu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	
	fmt.Println("Program started at:", time.Now().Format("15:04:05.000"))
	
	// Start a writer first
	wg.Add(1)
	go writeCounter(&wg, 100)
	time.Sleep(100 * time.Millisecond) // Give writer a head start

	// Start multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readCounter(&wg, i)
	}

	// Start another writer
	wg.Add(1)
	go writeCounter(&wg, 200)
	time.Sleep(100 * time.Millisecond) // Give writer a head start

	// Start more readers
	for i := 5; i < 10; i++ {
		wg.Add(1)
		go readCounter(&wg, i)
	}

	wg.Wait()
	fmt.Println("Program finished at:", time.Now().Format("15:04:05.000"))
}
