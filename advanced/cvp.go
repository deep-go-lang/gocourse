package advanced

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func PrintNumbers() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(time.Now().Format("15:04:05"))
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func PrintLetters() {
// 	for _, letter := range "Urvashi" {
// 		fmt.Println(time.Now().Format("15:04:05"))
// 		fmt.Println(string(letter))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
// func main() {
// 	go PrintNumbers()
// 	go PrintLetters()
// 	time.Sleep(5 * time.Second)
// }

func HeavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
	for i := 0; i < 100_000_000; i++ {
		
	}
	fmt.Println(time.Now().Format("15:04:05.000"))
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	numThreads := 10
	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup
	
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go HeavyTask(i, &wg)
	}
	wg.Wait()
	
}