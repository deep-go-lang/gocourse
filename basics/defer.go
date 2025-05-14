package basics

import "fmt"
func main() {

	process(15)

}

func process(i int){
	defer fmt.Println("deferred i value", i)
	defer fmt.Println("Process completed 2")
	defer fmt.Println("Process completed 3")
	defer fmt.Println("Process completed 4")
	i++
	fmt.Println("Processing...")
	fmt.Println("i value", i)
	defer fmt.Println("chromium deferred i value", i)
}