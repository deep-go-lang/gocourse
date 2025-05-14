package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	time.Sleep(time.Second * 2)

// 	fmt.Println("Starting timer...")

// 	timer := time.NewTimer(time.Second * 2)

// 	fmt.Println("Timer created")

// 	stopped := timer.Stop()

// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}

// 	timer.Reset(time.Second * 2)

// 	fmt.Println("Timer reset")

// 	<-timer.C

// 	fmt.Println("Timer expired")

// }

// func longRunningOperation() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {

// 	timeout := time.After(time.Second * 5)

// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	case <-timeout:
// 		fmt.Println("Operation timed out")
// 	}
// }

// func main() {
// 	timer := time.NewTimer(time.Second * 2)

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation completed")
// 	}()

	

// 	fmt.Println("Waiting...")

// 	time.Sleep(time.Second * 3)

// 	fmt.Println("Done")

	
	
// }


func main() {
	timer1 := time.NewTimer(time.Second * 1)
	timer2 := time.NewTimer(time.Second * 2)

	go func() {
		for {
			select {
			case <-timer1.C:
				fmt.Println("Timer 1 expired")
		    case <-timer2.C:
			fmt.Println("Timer 2 expired")
		}
	}
	}()

	time.Sleep(time.Second * 3)

	fmt.Println("Done")
}
