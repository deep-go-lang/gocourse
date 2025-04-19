package advanced

import (
	"fmt"
	"time"
)

func unbuffered_channels() {

	ch := make(chan int)

	go func() {
		ch <- 789
		time.Sleep(time.Second *2)
	}()

	time.Sleep(time.Second * 4)

	receiver := <-ch

	fmt.Println(receiver)

	

	

	

}
