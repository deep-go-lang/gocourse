package intermediate

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data := []byte("Hello, World!\n\n\n")

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully")

	file1, err := os.Create("test1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file1.Close()

	_, err = file1.WriteString("Hello, Urvashi!\n\n\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File1 written successfully")


	
}


