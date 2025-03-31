package intermediate

import "fmt"
func main() {

	i := 106_75.357_55

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Println("--------------------------------")

	string := "Hello, Urvashi!"

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)


	fmt.Println("--------------------------------")


	int := 15

	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%#X\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	fmt.Println("--------------------------------")
	

	txt := "Kriti!"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("%X\n", txt)
	fmt.Printf("% x\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)

	fmt.Println("--------------------------------")

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", f)

	fmt.Println("--------------------------------")

	flt := 4.8

	fmt.Printf("%f\n", flt)
	fmt.Printf("%e\n", flt)
	fmt.Printf("%g\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%-6.2f\n", flt)

	fmt.Println("--------------------------------")
	



	
	
	
	

}
