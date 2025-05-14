package advanced

import (
	"fmt"
	"strconv"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(time.Second * 2)
// 		done <- struct{}{}

// 	}()

// 	<-done

// 	fmt.Println("Finished...")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second * 1)
// 		ch <- 8
// 		fmt.Println("Sent value...")
// 	}()

// 	val := <-ch
// 	fmt.Println(val)

// func main() {

// 	num := 3
// 	ch := make(chan int, num)

// 	for i :=range num {
// 		time.Sleep(time.Second * 2)
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			ch <- id

// 		}(i)
// 	}

// 	for range num {
// 		fmt.Println("--------")
// 		data := <-ch
// 		fmt.Println("Data:", data)
// 	}

// 	fmt.Println("All Goroutines are finished!")

// }


func main() {
	data := make(chan string)

	go func() {
		for i := 0; i < 5; i++  {
			data <- "hello " + strconv.Itoa(i)
			time.Sleep(time.Millisecond * 100)

		}

		close(data)

	}()


	for value := range data {
		fmt.Println("Received value: ", value, ":", time.Now())
	}
}


	



