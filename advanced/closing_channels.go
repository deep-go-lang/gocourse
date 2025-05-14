package advanced

import "fmt"

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for value := range ch {
// 		fmt.Println(value)
// 	}

// }

// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	value, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is closed")
// 	} else {
// 		fmt.Println("Channel is open", value)
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for value := range ch {
// 		fmt.Println("Received",value)
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		close(ch)
// 		// close(ch)
// 	}()

// 	time.Sleep(time.Second * 1)

// }

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func filter (in <-chan int, out chan<- int) {
	for value := range in {
		if value % 2 == 0 {
			out <- value
		}
	}

	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for value := range ch2 {
		fmt.Println("Received", value)
	}
	
}


