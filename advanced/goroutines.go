package advanced

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("beginning of the program")
	go sayHello()
	fmt.Println("after sayHello")
	time.Sleep(2 * time.Second)

	fmt.Println("--------------------------------")

	go printNumbers()
	time.Sleep(5 * time.Second)
	go printLetters()
	time.Sleep(5 * time.Second)

	fmt.Println("--------------------------------")

	var err error

	go func() {
		err = doWork()
	}()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Work done successfully")
	}
	
	

	fmt.Println("end of the program")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello")
}


func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println("Numbers:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "Hello" {
		fmt.Println("Letters:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Some error")
}




