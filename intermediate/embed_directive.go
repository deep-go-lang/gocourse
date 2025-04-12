package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var exampleFile string

//go:embed basics
var basics embed.FS

func main() {
	fmt.Println(exampleFile)

	content, err := basics.ReadFile("basics/hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	fmt.Println("--------------------------------")

	err = fs.WalkDir(basics, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, d.Name(), d.Type())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}




