package advanced

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}


func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final value: %d\n", counter)
}

