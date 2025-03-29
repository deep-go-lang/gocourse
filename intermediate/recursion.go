package intermediate

import "fmt"

func main() {

	fmt.Println(factorial(5))

	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(12345))

	fmt.Println(sumOfDigits(123456789))

	fmt.Println(sumOfDigits(7))

	fmt.Println(fibonacci(10))



}


func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

