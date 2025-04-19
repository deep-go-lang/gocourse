package advanced

import (
	"fmt"
	"time"
)

func channels() {

	greeting := make(chan string)

	go func() {
		greeting <- "Hello"
		greeting <- "Urvashi"
		for _, v := range "Aishwarya" {
			greeting <- "alphabet" + string(v)
		}
	}()
	
	// go func() {
	// 	msg := <-greeting
	// 	fmt.Println(msg)
	// }()

	// go func() {
	// 	msg := <-greeting
	// 	fmt.Println(msg)
	// }()

	msg := <-greeting
	fmt.Println(msg)

	msg = <-greeting
	fmt.Println(msg)

	for i := 0; i < 9; i++ {
		msg := <-greeting
		fmt.Println(msg)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("end of program")
}
