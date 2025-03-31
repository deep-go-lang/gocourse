package main

import "fmt"

func main() {
	person1 := Person{
	firstName: "John", 
	lastName: "Doe", 
	age: 20, 
	address: Address{country: "India", city: "Mumbai", state: "Maharashtra"}, 
	PhoneNumber: PhoneNumber{Home: "75645553", mobile: "7544575667"}}

	person2 := Person{firstName: "Jane", age: 21, address: Address{country: "India", city: "Delhi", state: "Delhi"}}

	person3 := Person{firstName: "Jim", lastName: "Beam", address: Address{country: "USA", city: "New York", state: "New York"}}

	person4 := Person{firstName: "Jane", age: 21, address: Address{country: "India", city: "Delhi", state: "Delhi"}}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)
	fmt.Println(person1.age)
	fmt.Println(person1.PhoneNumber.mobile)
	fmt.Println(person1.Home)
	fmt.Println(person1.mobile)

	fmt.Println("person1 eauals to person2", person1 == person2)

	fmt.Println("person2 eauals to person4", person2 == person4)


	fmt.Println(person2.firstName)
	fmt.Println(person2.lastName)
	fmt.Println(person2.age)

	fmt.Println(person3.firstName)
	fmt.Println(person3.lastName)
	fmt.Println(person3.age)

	fmt.Println(person1.fullName())
	fmt.Println(person2.fullName())
	fmt.Println(person3.fullName())

	person1.incrementAgeByOne()
	fmt.Println(person1)

	fmt.Println("--------------------------------")

	actress := struct {
		firstName string
		lastName  string
		age       int
		email     string
		address   Address
	}{
		firstName: "Aishwarya",
		lastName:  "Rai",
		age:       40,
		email:     "aishwarya@gmail.com",
		address:   Address{country: "India", city: "Mumbai", state: "Maharashtra"},
	}

	fmt.Println(actress)

	actress.age = 41

	fmt.Println(actress)

}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneNumber
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
	fmt.Println("Age incremented by 1:", p.age)
}

type Address struct {
	country string
	city   string
	state  string
}

type PhoneNumber struct{
	Home string
	mobile string
}




