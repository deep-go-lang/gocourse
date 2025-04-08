package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		fmt.Println("Closing file")
		file.Close()
	}()

	fmt.Println("File opened successfully")

	// data := make([]byte, 1024)

	// count, err := file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Read", count, "bytes")
	// fmt.Println(string(data))

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
