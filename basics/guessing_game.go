package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())

	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Guess the number between 1 and 100")

	var guess int

	for {

		fmt.Println("Enter your guess: ")

		fmt.Scan(&guess)

		if guess == target {
			fmt.Println("You guessed the number!")
			break
		} else if guess < target {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")
		}
	}

}
