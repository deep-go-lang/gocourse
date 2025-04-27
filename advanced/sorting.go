package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     By
}

func (p *personSorter) Len() int { return len(p.people) }

func (p *personSorter) Swap(i, j int) { p.people[i], p.people[j] = p.people[j], p.people[i] }

func (p *personSorter) Less(i, j int) bool { return p.by(&p.people[i], &p.people[j]) }

func (by By) Sort(people []Person) {
	ps := &personSorter{people, by}
	sort.Sort(ps)
}

// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func (a ByName) Len() int           { return len(a) }
// func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {

	// numbers := []int{5, 2, 9, 1, 3, 6, 8, 4, 7}

	// sort.Ints(numbers)

	// fmt.Println(numbers)

	// strings := []string{"Urvashi", "Aishwarya", "Kriti", "Anushka", "Alia"}

	// sort.Strings(strings)

	// fmt.Println(strings)

	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Jim", Age: 35},
	}

	fmt.Println(people)

	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}

	By(ageAsc).Sort(people)

	fmt.Println("After sorting by age ascending: ", people)

	nameAsc := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	By(nameAsc).Sort(people)

	fmt.Println("After sorting by name ascending: ", people)

	fmt.Println("--------------------------------")

	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}

	By(ageDesc).Sort(people)

	fmt.Println("After sorting by age descending: ", people)

	nameDesc := func(p1, p2 *Person) bool {
		return p1.Name > p2.Name
	}

	By(nameDesc).Sort(people)

	fmt.Println("After sorting by name descending: ", people)

	fmt.Println("--------------------------------")

	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(lenName).Sort(people)

	fmt.Println("After sorting by length of name: ", people)

	fmt.Println("--------------------------------")

	stringSlice := []string{"Urvashi", "Aishwarya", "Kriti", "Anushka", "Alia"}

	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})

	fmt.Println("After sorting by last character: ", stringSlice)

	// sort.Sort(ByAge(people))

	// fmt.Println(people)

	// fmt.Println("--------------------------------")

	// sort.Sort(ByName(people))

	// fmt.Println(people)
}
