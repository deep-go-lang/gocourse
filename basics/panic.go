package basics

import "fmt"

func main() {
	// process(10)
	process(-10)

}

func process(i int){

	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")
	defer fmt.Println("deferred 3")

	if i < 0 {
		fmt.Println("This state is before panic")
		panic("i cannot be negative")
		// fmt.Println("This state is after panic")
		// defer fmt.Println("deferred 4")
	}
	fmt.Println("i value", i)
}

