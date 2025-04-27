package advanced

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	pool := sync.Pool{
		// New: func() interface{} {
		// 	fmt.Println("Creating new Person")
		// 	return &Person{}
		// },
	}

	pool.Put(&Person{Name: "Aishwarya", Age: 40})

	person1 := pool.Get().(*Person)
	person1.Name = "John"
	person1.Age = 20
	fmt.Println("Person 1:", person1)
	fmt.Printf("Person 1: %+v\nName: %s\nAge: %d\n", person1, person1.Name, person1.Age)

	pool.Put(person1)

	fmt.Println("Returned person1 to pool")

	fmt.Println("\n--------------------------------")

	person2 := pool.Get().(*Person)
	fmt.Println("Person 2:", person2)
	fmt.Printf("Person 2: %+v\nName: %s\nAge: %d\n", person2, person2.Name, person2.Age)

	pool.Put(person2)

	fmt.Println("Returned person2 to pool")

	fmt.Println("\n--------------------------------")

	person3 := pool.Get().(*Person)
	fmt.Println("Person 3:", person3)
	fmt.Printf("Person 3: %+v\nName: %s\nAge: %d\n", person3, person3.Name, person3.Age)

	pool.Put(person3)

	fmt.Println("Returned person3 to pool")

	fmt.Println("\n--------------------------------")

	person4 := pool.Get().(*Person)
	fmt.Println("Person 4:", person4)
	fmt.Printf("Person 4: %+v\nName: %s\nAge: %d\n", person4, person4.Name, person4.Age)

	pool.Put(&Person{Name: "Namrata", Age: 45})

	person5 := pool.Get().(*Person)
	if person5 != nil {
		fmt.Println("Person 5:", person5)
		person5.Name = "Neha"
		person5.Age = 35
		fmt.Printf("Person 5: %+v\nName: %s\nAge: %d\n", person5, person5.Name, person5.Age)
	} else {
		fmt.Println("Person 5 is nil")
	}

	fmt.Println("\n--------------------------------")

	pool.Put(person4)
	pool.Put(person5)

	fmt.Println("Returned person4 and person5 to pool")

	person6 := pool.Get().(*Person)
	fmt.Println("Person 6:", person6)
	fmt.Printf("Person 6: %+v\nName: %s\nAge: %d\n", person6, person6.Name, person6.Age)

	person7 := pool.Get().(*Person)
	fmt.Println("Person 7:", person7)
	fmt.Printf("Person 7: %+v\nName: %s\nAge: %d\n", person7, person7.Name, person7.Age)

	person7.Name = "Urvashi"
	person7.Age = 21

	pool.Put(person7)

	fmt.Println("Returned person7 to pool")
	
	
	person8 := pool.Get().(*Person)
	fmt.Println("Person 8:", person8)
	fmt.Printf("Person 8: %+v\nName: %s\nAge: %d\n", person8, person8.Name, person8.Age)

	

	
	
	
	
	
	
}

