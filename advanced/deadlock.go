package advanced

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("goroutine 1 locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("goroutine 1 finished")
	}()

	go func() {
		mu1.Lock()
		fmt.Println("goroutine 2 locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("goroutine 2 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("goroutine 2 finished")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("main goroutine finished")
	// select {}

}
