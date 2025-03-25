package basics

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 4; i <= 14; i++ {
	// 	fmt.Println(i)
	// }

	// number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := 0; i < len(number); i++ {
	// 	fmt.Println(number[i])
	// }

	// for index, value := range number {
	// 	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// for _, value := range number {
	// 	fmt.Println(value)
	// }

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd number: ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10 {
		i++
		fmt.Println(10 - i)
	}

	fmt.Println("Done")
}
