package intermediate

import "fmt"

type rectangle struct {
	width float64
	height float64
}

type shape struct {
	rectangle
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

type myInt int

func (m myInt) isPositive() bool {
	return m > 0
}

func (m *myInt) increment() {
	*m++
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("Area of rectangle is:", r.area())

	r.scale(2)
	fmt.Println("Area of rectangle is:", r.area())

	fmt.Println("Rectangle is:", r)

	num1 := myInt(-5)
	num2 := myInt(10)

	fmt.Println("Is number positive?", num1.isPositive())
	fmt.Println("Is number positive?", num2.isPositive())

	num1.increment()
	fmt.Println("Number is:", num1)

	num2.increment()
	fmt.Println("Number is:", num2)

	s := shape{rectangle: rectangle{width: 10, height: 5}}
	fmt.Println("Area of shape is:", s.area())

	s.scale(2)
	fmt.Println("Area of shape is:", s.area())

}



