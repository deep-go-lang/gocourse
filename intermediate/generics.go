package intermediate

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	items []T
}

type newStack[T string] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *newStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

	func (s *Stack[T]) isEmpty() bool {
		return len(s.items) == 0
	}

	func (s *newStack[T]) isEmpty() bool {
		return len(s.items) == 0
	}

func (s *Stack[T]) size() int {
	return len(s.items)
}

func (s *Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	for _, item := range s.items {
		fmt.Print(item, " ")
		
	}	

	fmt.Println()
}

func (s *newStack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	for _, item := range s.items {
		fmt.Print(item, " ")
		
	}	

	fmt.Println()
}


func main() {

	// s := Stack[int]{}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.printAll()

	// fmt.Println("--------------------------------")
	// fmt.Println(s.Pop())
	// s.printAll()
	// fmt.Println(s.Pop())
	// s.printAll()
	// fmt.Println(s.Pop())
	// s.printAll()

	s := Stack[string]{}
	s.Push("Hello")
	s.Push("World")
	s.printAll()
	fmt.Println(s.Pop())
	s.printAll()
	s.Push("Aishwarya")
	s.printAll()

	fmt.Println("--------------------------------")

	ns := newStack[string]{}
	ns.Push("Hello")
	ns.Push("World")
	ns.printAll()
	fmt.Println(ns.isEmpty())

	

	
	
	
	
	

	
	

	// a, b := swap(1, 2)
	// fmt.Println(a, b)

	// x, y := swap("hello", "world")
	// fmt.Println(x, y)

	
	
}
