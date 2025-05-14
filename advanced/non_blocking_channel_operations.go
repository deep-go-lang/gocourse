package advanced

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)

	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received Message:", msg)
	// default:
	// 	fmt.Println("No Message Received")
	// }

	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent Message to Channel")
	// default:
	// 	fmt.Println("Channel is full")
	// }


	data := make(chan int)
	quit := make(chan bool)

	go func () {
		for {
			select {
				case d := <- data:
					fmt.Println("Received Data:", d)
				case <- quit:
					fmt.Println("Quitting...")
					return
				default:
					fmt.Println("Waiting for Data...")
					time.Sleep(time.Second * 2)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		data <- i
		time.Sleep(time.Second * 1)
	}
	quit <- true
	

}
