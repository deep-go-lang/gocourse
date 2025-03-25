package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 35, 16
	var result int

	result = a + b

	fmt.Println("Sum of a and b is:", result)

	result = a - b

	fmt.Println("Difference of a and b is:", result)

	result = a * b

	fmt.Println("Product of a and b is:", result)

	result = a / b	

	fmt.Println("Quotient of a and b is:", result)

	result = a % b

	fmt.Println("Remainder of a and b is:", result)

	const pi float64 = 22.0 / 7.0

	fmt.Println("Value of pi is:", pi)

	var maxInt int64 = 9223372036854775807

	fmt.Println("Max integer value is:", maxInt)

	maxInt = maxInt + 1

	fmt.Println("Max integer value after increment is:", maxInt)

	var uMaxInt uint64 = 18446744073709551615

	fmt.Println("Max unsigned integer value is:", uMaxInt)

	uMaxInt = uMaxInt + 1

	fmt.Println("Max unsigned integer value after increment is:", uMaxInt)

	var minFloat float64 = 1.0e-323

	fmt.Println("Min float value is:", minFloat)

	minFloat = minFloat / math.MaxFloat64

	fmt.Println("Min float value after division is:", minFloat)

	fmt.Println(math.MinInt64)
	

}

