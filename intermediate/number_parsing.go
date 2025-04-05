package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	number := "1234567890"

	parsedNumber, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error parsing number:", err)
	}

	fmt.Println("Parsed number:", parsedNumber)
	fmt.Println(parsedNumber + 1)

	parsedNumber2, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		fmt.Println("Error parsing number:", err)
	}

	fmt.Println("Parsed number:", parsedNumber2)
	fmt.Println(parsedNumber2 + 1)

	floatString := "123.456"
	parsedFloat, err := strconv.ParseFloat(floatString, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
	}

	fmt.Printf("Parsed float: %.2f\n", parsedFloat)
	fmt.Printf("Parsed float: %.2f\n", parsedFloat+1)

	binaryString := "1010"
	parsedBinary, err := strconv.ParseInt(binaryString, 2, 32)
	if err != nil {
		fmt.Println("Error parsing binary:", err)
	}

	fmt.Println("Parsed binary:", parsedBinary)
	fmt.Println(parsedBinary + 1)

	hexString := "1A"
	parsedHex, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		fmt.Println("Error parsing hex:", err)
	}

	fmt.Println("Parsed hex:", parsedHex)
	fmt.Println(parsedHex + 1)

	number1 := "1234abc"

	parsedNumber1, err := strconv.Atoi(number1)
	if err != nil {
		fmt.Println("Error parsing number:", err)
	}

	fmt.Println("Parsed number:", parsedNumber1)
	fmt.Println(parsedNumber1 + 1)

	
	
}


