package basics


import "fmt"
func main() {

	// fruit := "mango"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("Apple")
	// case "banana":
	// 	fmt.Println("Banana")
	// default:
	// 	fmt.Println("Unknown fruit")
	// }


	// day := "wednesday"

	// switch day {
	// case "monday", "tuesday", "wednesday", "thursday", "friday":
	// 	fmt.Println("It is a weekday")
	// case "saturday", "sunday":
	// 	fmt.Println("It is a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 25


	// switch {

	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 20")
	// case number >= 20 && number < 30:
	// 	fmt.Println("Number is between 20 and 30")
	// default:
	// 	fmt.Println("Number is equal to or greater than 30")
	// }


	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Number is greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is equal to 2")
	// 	fallthrough
	// case num == 3:
	// 	fmt.Println("Number is equal to 3")
	// default:
	// 	fmt.Println("Number is greater than 2")
	// }

	checkType(4)
	checkType(4.5)
	checkType("Hello")
	checkType(true)


	score := 66

	switch {
	case score >= 80:
		fmt.Println("A")
	case score >= 70:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}
}


func checkType (x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case int32:
		fmt.Println("x is an int32")
	case float64:
		fmt.Println("x is a float")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of an unknown type")
	}
}
