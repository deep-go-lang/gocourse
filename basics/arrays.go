package basics


import "fmt"

func main() {

	var numbers [5]int

	fmt.Println(numbers)

	numbers[4] = 578

	fmt.Println(numbers)

	numbers[0] = 57564

	fmt.Println(numbers)

	fruits := [4]string{"apple", "banana", "cherry", "orange"}

	fmt.Println("Fruits:", fruits)

	fmt.Println("Fruit at index 3:", fruits[3])

	originalArray := [5]int{1, 2, 3, 4, 5}

	copiedArray := originalArray

	copiedArray[3] = 500

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index", i, ":", fruits[i])
	}

	for index, fruit := range fruits {
		fmt.Println("Fruit at index", index, ":", fruit)
	}

	for _, v := range fruits {
		fmt.Printf("Fruit: %v\n", v)
	}

	a, _ := someFunction()

	fmt.Println(a)

	b := 4

	_ = b

	fmt.Println("The length of Fruits array is:", len(fruits))

	array1 := [3]int{1, 5, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("array1 == array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		 {4, 5, 6}, 
		 {7, 8, 9},
		}

	fmt.Println("Matrix:", matrix)

	newOriginalArray := [3]int{1, 2, 3}

	var newCopiedArray *[3]int

	newCopiedArray = &newOriginalArray

	newCopiedArray[2] = 100

	fmt.Println("New original array:", newOriginalArray)
	fmt.Println("New copied array:", *newCopiedArray)


	
	
	
	
	
	
	
}

func someFunction() (int, int) {
	return 1, 2
}
