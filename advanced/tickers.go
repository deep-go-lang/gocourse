package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	ticker := time.NewTicker(time.Second * 2)

// for t := range ticker.C {
// 	fmt.Println("Tick at", t)
// }

// i := 1
// for range ticker.C {
// 	i *= 2
// 	fmt.Println(i)
// 	if i > 10 {
// 		ticker.Stop()
// 		break
// 	}
// }

// }

// func periodicTask() {
// 	fmt.Println("Periodic task", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second * 2)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {
// 	ticker := time.NewTicker(time.Second * 2)
// 	stop := time.After(time.Second * 10)

// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-stop:
// 			fmt.Println("Stopping ticker")
// 			return
// 		case t := <-ticker.C:
// 			fmt.Println("Tick at", t)
// 		}
// 	}
// }

func main() {
	// Create multiple tickers with different intervals
	fastTicker := time.NewTicker(time.Second * 1)    // Ticks every 1 second
	mediumTicker := time.NewTicker(time.Second * 2)  // Ticks every 2 seconds
	slowTicker := time.NewTicker(time.Second * 3)    // Ticks every 3 seconds

	// Create a stop channel to terminate the program after 10 seconds
	stop := time.After(time.Second * 10)

	// Make sure to stop all tickers when done
	defer func() {
		fastTicker.Stop()
		mediumTicker.Stop()
		slowTicker.Stop()
	}()

	fmt.Println("Starting multiple tickers...")

	for {
		select {
		case <-stop:
			fmt.Println("Stopping all tickers")
			return
		case t := <-fastTicker.C:
			fmt.Println("Fast ticker:", t.Format("15:04:05"))
		case t := <-mediumTicker.C:
			fmt.Println("Medium ticker:", t.Format("15:04:05"))
		case t := <-slowTicker.C:
			fmt.Println("Slow ticker:", t.Format("15:04:05"))
		}
	}
}
