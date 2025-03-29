package basics

import "fmt"

import "os"

func main() {
	defer fmt.Println("deferred line")
	fmt.Println("Starting program")
	os.Exit(0)
	fmt.Println("Ending program")

}
