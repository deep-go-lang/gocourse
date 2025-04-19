package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second * 1)

// 		ch1 <- 1

// 	}()

// 	go func() {
// 		// time.Sleep(time.Second * 1)

// 		ch2 <- 2

// 	}()

// 	time.Sleep(time.Second * 2)

// 	for range 2 {

// 		select {

// 		case msg := <-ch1:
// 			fmt.Println("Received from channel 1:", msg)

// 		case msg := <-ch2:
// 			fmt.Println("Received from channel 2:", msg)

// 		default:
// 			fmt.Println("No channels are ready...")

// 		}

// 	}

// 	fmt.Println("End of Program")

// }

// func main () {
// 	ch := make(chan int)

// 	go func(){
// 		time.Sleep(time.Second * 2)
// 		ch <- 2
// 		close(ch)
// 	}()

// 	select {
// 	case msg, ok := <- ch:
// 		if !ok {
// 			fmt.Println("Channel Closed...")
// 			return
// 		}
// 		fmt.Println("Received Message:", msg)

// 	case <- time.After(time.Second * 3):
// 		fmt.Println("Timeout")

		
// 	}
// }
func main () {
	ch := make(chan int)

	go func () {
		ch <- 1
		time.Sleep(time.Second * 4)
		close(ch)
	}()

	for {
		select {
		case msg, ok := <- ch:
			if !ok {
				fmt.Println("Channel Closed...")
				return
			}
			fmt.Println("Received Message:", msg)
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout")
		}
	}
}
