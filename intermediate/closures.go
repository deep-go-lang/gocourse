package intermediate

import "fmt"

func main() {

	// sequence1 := adder()

	// fmt.Println(sequence1())
	// fmt.Println(sequence1())
	// fmt.Println(sequence1())
	// fmt.Println(sequence1())

	// fmt.Println("--------------------------------")

	// sequence2 := adder()

	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())

	subtracter := func() func(int) int {

		countdown := 99

		return func(x int) int {

			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}


// func adder() func() int {

// 	i := 0

// 	fmt.Println("Previous value of i:", i)

// 	return func() int {
// 		i++
// 		fmt.Println("Adding 1 to i:")
// 		return i
// 	}

// }
