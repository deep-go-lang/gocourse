package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	employeeInfo person
	salary float64
	empId  int
}

func (e employee) display() {
	fmt.Println(e.employeeInfo.name, e.employeeInfo.age, e.salary, e.empId)
}

func (p person) introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old\n", p.name, p.age)
}

func (e employee) introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old and my salary is %.2f and my employee id is %d\n", e.employeeInfo.name, e.employeeInfo.age, e.salary, e.empId)
}

func main() {
	e := employee{employeeInfo: person{name: "John", age: 30}, salary: 50000, empId: 1}
	e.display()

	p := person{name: "Jane", age: 25}
	p.introduce()

	e.introduce()
}
