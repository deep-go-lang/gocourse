package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	

	// reader := strings.NewReader("Hello, World!\nThis is a test for bufio\nnot for strings")

	// buf := bufio.NewReader(reader)

	// data := make([]byte, 20)

	// n, err := 	buf.Read(data)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(data[:n]))

	// fmt.Println("--------------------------------")

	// line, err := buf.ReadString('\n')

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(line)

	writer := bufio.NewWriter(os.Stdout)

	n, err := 	writer.WriteString("Hello, World!\n")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	writer.Flush()

	fmt.Println(n)

	fmt.Println("--------------------------------")

	data := []byte("Hello, Urvashi!\n")

	n1, err := writer.Write(data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	err = writer.Flush()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(n1)

	fmt.Println("--------------------------------")

	stringToWrite := "Hello, Urvashi!\nYou are a good girl!\n"

	n2, err := writer.WriteString(stringToWrite)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	err = writer.Flush()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println(n2)

}
