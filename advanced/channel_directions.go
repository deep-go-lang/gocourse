package advanced

import "fmt"

func main() {

	ch := make(chan  int)

	producer(ch)

	consumer(ch)

}

func producer(ch chan <- int) {
	go func() {

		for i := 0; i < 5; i++ {
			ch <- i
		}

		close(ch)

	}()
}

func consumer(ch <- chan int) {

	for value := range ch {
		fmt.Println("Received: ", value)
	}

}
