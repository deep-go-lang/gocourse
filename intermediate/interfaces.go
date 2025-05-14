package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type rectangle2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r rectangle2) area() float64 {
	return r.width * r.height
}

func (r rectangle2) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	r2 := rectangle2{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r2)

	// fmt.Println(r.area())
	// fmt.Println(c.perimeter())
	// fmt.Println(c.area())
	// fmt.Println(r.perimeter())
	// fmt.Println(c.diameter())

	myPrinter("Hello", 1, 56.7, true)

	printType(1)
	printType("Hello")
	printType(56.7)
	printType(true)
}

func myPrinter(i... interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown type")
	}
}


