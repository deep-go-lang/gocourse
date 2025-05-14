package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	err := processData2()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data processed successfully")

	
	
}

type customError struct {
	message string
	code int
	err error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %s\n", e.code, e.message, e.err)
}

// func processData() error {
// 	return &customError{message: "Data processing failed", code: 500}
// }

func processData() error {
	return  errors.New("Internal server error")}


func processData2() error {
	err := processData()
	if err != nil {
		return &customError{message: "Data processing failed", code: 500, err: err}
	}
	return nil
}



