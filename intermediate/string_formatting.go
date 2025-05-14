package intermediate

import "fmt"

func main() {

	num := 1235

	fmt.Printf("Number: %05d\n", num)

	message := "Hello"

	fmt.Printf("|%10s|\n", message)

	fmt.Printf("|%-10s|\n", message)

	message2 := "Hello,\nWorld!"

	message3 := `Hello,\nWorld!`

	fmt.Printf("%s\n", message2)

	// fmt.Printf("%q\n", message3)

	fmt.Println(message3)

	
	
	
	
	
	
	

}
