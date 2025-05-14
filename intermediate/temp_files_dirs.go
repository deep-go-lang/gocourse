package intermediate

import (
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create and manage temp file
	tempFile, err := os.CreateTemp("", "tempFile.txt")
	checkError(err)
	fmt.Println("Temp file created:", tempFile.Name())
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write some content to verify file operations
	_, err = tempFile.WriteString("Hello, World!")
	checkError(err)
	fmt.Println("Content written to temp file")

	fmt.Println("--------------------------------")

	// Create and manage temp directory
	tempDir, err := os.MkdirTemp("", "tempDir")
	checkError(err)
	fmt.Println("Temp directory created:", tempDir)
	defer os.RemoveAll(tempDir)
}

// func createTempFile() {
// 	tempFile, err := os.CreateTemp("", "tempFile.txt")
// 	checkError(err)
// 	defer os.Remove(tempFile.Name())
// }
