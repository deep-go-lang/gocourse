package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	joinedPath := filepath.Join(".", "data", "myfile.txt")
	fmt.Println("Joined Path:", joinedPath)

	absPath, err := filepath.Abs(joinedPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Absolute Path:", absPath)

	dir, file := filepath.Split(absPath)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
	

	cleanedPath := filepath.Clean(".\\data\\..\\data\\file.txt")
	fmt.Println("Cleaned Path:", cleanedPath)

	basePath := filepath.Base(absPath)
	fmt.Println("Base Path:", basePath)

	ext := filepath.Ext(absPath)
	fmt.Println("Extension:", ext)

	isAbs := filepath.IsAbs(joinedPath)
	fmt.Println("Is Absolute:", isAbs)

	trimmedPath := strings.TrimSuffix(absPath, ext)
	fmt.Println("Trimmed Path:", trimmedPath)

	// Get the current working directory as base path
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	fmt.Println("Base Directory:", baseDir)

	relPath, err := filepath.Rel(baseDir, absPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Relative Path:", relPath)


	relPath, err = filepath.Rel("a/b/c/file.png", "a/b/f/myfile.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Relative Path:", relPath)

	
	

	
	
}
