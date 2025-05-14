package advanced

import (
	"fmt"
	"sync"
)

var once sync.Once

func Initialize() {
	fmt.Println("This will only run once")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine", i)
			once.Do(Initialize)
			// Initialize()
		}()
	}
	wg.Wait()
}
