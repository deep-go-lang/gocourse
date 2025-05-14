package intermediate

import "fmt"

func main() {

	fmt.Print("Hello, World!")
	fmt.Print(12, 500)

	fmt.Println("--------------------------------")

	fmt.Println("Hello, World!")
	fmt.Println(12, 500)

	fmt.Println("--------------------------------")

	name := "Urvashi"
	age := 20

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	fmt.Println("--------------------------------")

	s := fmt.Sprintf("Name: %s, Age: %d\n", name, age)
	fmt.Print(s)

	v := fmt.Sprint("Urvashi", "Rautela", 356, 798)
	fmt.Print(v)

	k := fmt.Sprintln("Urvashi", "Rautela", 356, 798)
	fmt.Print(k)
	fmt.Print(k)

	fmt.Println("--------------------------------")

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)


	// fmt.Println("--------------------------------")

	// var nm string
	// var ag int
	// fmt.Print("Enter your name and age: ")
    // fmt.Scan(&nm, &ag)
	// fmt.Printf("Name: %s, Age: %d\n", nm, ag)

	// fmt.Println("--------------------------------")

	// var fruit string
	// var price float64

	// fmt.Print("Enter your fruit and price: ")
	// fmt.Scanln(&fruit, &price)
	// fmt.Printf("Fruit: %s, Price: %f\n", fruit, price)

	fmt.Println("--------------------------------")

	// var country string
	// var capital string
	// var population int

	// fmt.Print("Enter your country, capital and population: ")
	// fmt.Scanf("%s %s %d", &country, &capital, &population)
	// fmt.Printf("Country: %s, Capital: %s, Population: %d\n", country, capital, population)

	err := checkAge(16)
	if err != nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You may enter the club")
	}



}

func checkAge(age int) error{
	if age < 18{
		return fmt.Errorf("age %d is less than 18", age)
	}
	return nil
}


