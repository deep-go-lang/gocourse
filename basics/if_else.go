package basics


import "fmt"

func main() {

	// age := 11

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else if age >= 13 {
	// 	fmt.Println("You are a teenager")
	// } else {
	// 	fmt.Println("You are a child")
	// }

	// temperature := 25

	// if temperature > 30 {
	// 	fmt.Println("It is hot")
	// } else if temperature <= 30 && temperature >= 20 {
	// 	fmt.Println("It is warm")
	// } else {
	// 	fmt.Println("It is cold")
	// }

	// num := 18

	// if num%2 == 0 {
	// 	fmt.Println("Even")
	// 	if num%3 == 0 {
	// 		fmt.Println("Divisible by 3")
	// 	} else {
	// 		fmt.Println("Not divisible by 3")
	// 	}
	// } else {
	// 	fmt.Println("Odd")
	// }

	if 15%2 == 0 || 17%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	num := 56

	if num%2 == 0 && num%5 == 0 {
		fmt.Println("Divisible by 2 and 5")
	} else {
		fmt.Println("Not divisible by 2 and 5")
	}

}
