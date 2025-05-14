package intermediate

import "fmt"

func main() {

	var a int = 10
	var b *int = &a
	var c *int
	fmt.Println(c)



	fmt.Println(a, b)
	fmt.Println(*b)
	fmt.Println(&b)
	fmt.Println(&a)
	fmt.Println(&*b)
	fmt.Println(*&a)
	fmt.Println(**&b)

	modifyValue(b)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
