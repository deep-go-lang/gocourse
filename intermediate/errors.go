package intermediate

import (
	"errors"
	"fmt"
	"math"
)

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("negative number")
	}
	return math.Sqrt(n), nil
}

func processData(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: empty data")
	}
	return nil
}

func main() {

	// sqrt, err := squareRoot(-16)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(sqrt)
	// }

	data := []byte{1, 2, 3}
	fmt.Println(data)
	err := processData(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data processed successfully")
	}

	// err1 := eProcessData()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// } 

	// b := 	err1.Error()
	// fmt.Println(b)

	// err1 := eProcessData()
	// fmt.Println(err1)

	// newStruct := myError{message: "My Data"}

	// fmt.Println(newStruct.Error())

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data read successfully")

	

	
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func eProcessData() error {
	return &myError{"custom error message"}
}

func readConfig() error {
	return errors.New("config error")
}


func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

