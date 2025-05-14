package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	const MAXTRIALS = 3

	var employeeID = 56864

	fmt.Println("Employee ID:", employeeID)

}
