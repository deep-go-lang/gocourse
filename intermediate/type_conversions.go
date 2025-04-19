package intermediate

import "fmt"

func main() {

	var a int = 20
	b := int32(a)
	c := float64(b)

	fmt.Println(a, b, c)

	fmt.Println("--------------------------------")

	e := 3.14565

	f := int(e)

	fmt.Println(e, f)

	fmt.Println("--------------------------------")

	g := "Urvashi Rautela❤❤💖💗💗"

	h := []byte(g)

	fmt.Println(g, h)

	fmt.Println("--------------------------------")

	i := []rune(g)

	fmt.Println(i)

	fmt.Println("--------------------------------")

	j := []byte{255, 76, 100}

	k := string(j)

	fmt.Println(j, k)

	fmt.Println("--------------------------------")

	l := []rune{255, 76, 100}

	m := string(l)

	fmt.Println(l, m)

	fmt.Println("--------------------------------")
	

	char := []byte{'a', 'b', 'c'}

	n := string(char)

	fmt.Println(char, n)

	fmt.Println("--------------------------------")

	
	
	
	
}
