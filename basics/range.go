package basics

import "fmt"

func main() {

	name := "Aishwarya Rai"

	for i, v := range name {
		fmt.Printf("Index: %d, Value: %c\n", i, v)
		fmt.Println("Index:", i, "Value:", v)
		fmt.Println("--------------------------------")
	}



}
