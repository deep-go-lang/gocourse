package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(val.Intn(1001))

// 	fmt.Println(rand.Intn(1001) )

// 	fmt.Println(rand.Float64())

	for {
		fmt.Println("Welcome to the dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int

		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid input, please enter 1 or 2")
			continue
		}

		if choice == 1 {
			die1 := rand.Intn(6) + 1
			die2 := rand.Intn(6) + 1
			fmt.Printf("You rolled a %d and a %d\n", die1, die2)
			fmt.Printf("Total: %d\n", die1+die2)

			fmt.Print("Do you want to roll again? (y/n): ")
			var rollAgain string
			_, err := fmt.Scan(&rollAgain)
			if err != nil || (rollAgain != "y" && rollAgain != "n") {
				fmt.Println("Invalid input, please enter y or n")
				continue
			}
			if rollAgain == "n" {
				fmt.Println("Thanks for playing!")
				break
			}
			fmt.Println("Rolling again...")
			time.Sleep(1 * time.Second)
			continue
		} else if choice == 2 {
			fmt.Println("Exiting the game")
			break
		}


	}
}
