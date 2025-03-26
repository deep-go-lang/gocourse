package basics

import "fmt"

func main() {

	// i := 0

	// for i < 10 {
	// 	fmt.Println("Iteration: ", i)
	// 	i++
	// }
	
	// sum := 0

	// for {
	// 	sum += 10
	// 	fmt.Println("Sum: ", sum)
	// 	if sum >= 100 {
	// 		break
	// 	}
	// }

	num := 1

	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number: ", num)
		num++
	}
}
