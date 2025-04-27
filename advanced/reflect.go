package advanced

import (
	"fmt"
	"reflect"
)

// type User struct {
// 	Name string
// 	Age int
// }

type Greeter struct{}

func (g Greeter) Hello(firstName string, lastName string) string {
	return "Hello, " + firstName + " " + lastName + "!"
}

func main() {
	g := Greeter{}

	v := reflect.ValueOf(g)

	fmt.Println("type of v:", v.Type())

	m := v.MethodByName("Hello")

	results := m.Call([]reflect.Value{reflect.ValueOf("Aishwarya"), reflect.ValueOf("Rai")})

	fmt.Println("results:", results[0].String())

	fmt.Println("method of v:", m)

	t := reflect.TypeOf(g)

	fmt.Println("type of t:", t)

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method: %d: %v\n", i, t.Method(i).Name)
	}

	
	
	
	
	
	
	
	
	
	
	
}

// func main() {

	// p := User{
	// 	Name: "Urvashi",
	// 	Age: 20,
	// }

	// v := reflect.ValueOf(p)

	// fmt.Println("type of v:", v.Type())

	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Printf("field: %d: %v\n", i, v.Field(i))
	// }

	// fmt.Println("--------------------------------")

	// v1 := reflect.ValueOf(&p).Elem()

	// nameField := v1.FieldByName("Name")

	// fmt.Println("nameField:", nameField)

	// ageField := v1.FieldByName("Age")

	// fmt.Println("ageField:", ageField)

	// if nameField.CanSet() {
	// 	nameField.SetString("Aishwarya")
	// }

	// fmt.Println("Modified nameField:", nameField)

	// fmt.Println("Modified person:", p)
	
	
	
// }

// func main() {

	// x := 10

	// v := reflect.ValueOf(x)

	// t := v.Type()

	// fmt.Println("value:", v)
	// fmt.Println("type:", t)
	// fmt.Println("kind:", v.Kind())
	// fmt.Println("Is string:", v.Kind() == reflect.String)
	// fmt.Println("is zero:", v.IsZero())

	// fmt.Println("--------------------------------")

	// y := 10

	// v2 := reflect.ValueOf(&y).Elem()

	// v3 := reflect.ValueOf(&y)

	// t2 := v2.Type()

	// t3 := v3.Type()

	// fmt.Println("value of v2:", v2)
	// fmt.Println("value of v3:", v3)
	// fmt.Println("type of v2:", t2)
	// fmt.Println("type of v3:", t3)
	// fmt.Println("kind of v2:", v2.Kind())
	// fmt.Println("kind of v3:", v3.Kind())

	// fmt.Println("Original value of v2", v2.Int())

	// v2.SetInt(20)

	// fmt.Println("Modified value of v2", v2.Int())

	// fmt.Println("--------------------------------")

	// var intf interface{} = "Hello, World!"

	// v4 := reflect.ValueOf(intf)

	// fmt.Println("type of v4:", v4.Type())

	// if v4.Kind() == reflect.String {
	// 	fmt.Println("v4 is a string")
	// }

	
	
// }
