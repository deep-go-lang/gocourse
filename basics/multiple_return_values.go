package basics

import "fmt"

import "errors"

func main() {

	q, r := division(15, 4)

	fmt.Printf("The quotient is %v\nand the remainder is %v\n", q, r)

	result, err := compare(10, 10)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}

func division(x int, y int) (quotient int, remainder int) {
    quotient = x / y
    remainder = x % y
    return 
}

func compare(x int, y int) (string, error) {
	if x > y {
		return "x is greater than y", nil
	} else if x < y {
		return "x is less than y", nil
	} else {
		return "", errors.New("x and y are equal")
	}
}

