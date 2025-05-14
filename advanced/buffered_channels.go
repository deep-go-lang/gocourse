package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		ch <- 4
// 		ch <- 9
// 	}()

// 	fmt.Println("value: ", <-ch)

// 	fmt.Println("value: ", <-ch)

// 	fmt.Println("End of Program")

// }

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println("Receiving from Buffer")

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("received", <-ch)
	}()

	fmt.Println("Blocking Starst")

	ch <- 3

	fmt.Println("Blocking Ends")

	// fmt.Println("received", <-ch)
	// fmt.Println("received", <-ch )

}
