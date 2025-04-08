package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {


	file, err := os.Open("filtered_lines.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	
	defer func() {
		fmt.Println("Closing file")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	keyword := "important"

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "vital")
			fmt.Printf("Line %d: %s\n", lineNumber, line)
			fmt.Printf("Filtered Line %d: %s\n", lineNumber, updatedLine)
			lineNumber++
			fmt.Println("--------------------------------")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	

	

}
