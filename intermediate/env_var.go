package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("user:", user)
	fmt.Println("home:", home)

	fmt.Println("--------------------")

	err := os.Setenv("ACTRESS", "Kriti Sanon")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}

	actress := os.Getenv("ACTRESS")
	fmt.Println("actress:", actress)

	err = os.Unsetenv("ACTRESS")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
	}

	fmt.Println("actress:", os.Getenv("ACTRESS"))

	fmt.Println("--------------------")

	for _, env := range os.Environ() {
		keyValuePair := strings.SplitN(env, "=", 2)
		fmt.Println(keyValuePair[0])
	}

	fmt.Println("--------------------")

	

}
