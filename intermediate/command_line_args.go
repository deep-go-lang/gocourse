package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("command", os.Args[0])

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	fmt.Println("--------------------------------")

	if len(os.Args) > 1 {
		fmt.Println("Argument1", os.Args[1])
	}

	fmt.Println("--------------------------------")

	if len(os.Args) > 2 {
		fmt.Println("Argument2", os.Args[2])
	}

	fmt.Println("--------------------------------")

	var name string
	var age int
	var female bool

	flag.StringVar(&name, "name", "Urvashi", "Name of the user")
	flag.IntVar(&age, "age", 20, "Age of the user")
	flag.BoolVar(&female, "female", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name", name)
	fmt.Println("Age", age)
	fmt.Println("Female", female)

	

}
