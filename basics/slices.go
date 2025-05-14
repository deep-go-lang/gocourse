package basics

import "fmt"

import "slices"

func main() {

	// var numbers []int

	// var numbers1 = []int{1, 2, 3, 4, 5}

	// numbers2 := []int{1, 2, 3, 4, 5}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}

	sicea := a[1:4]

	fmt.Println(sicea)

	sicea = append(sicea, 6, 45, 98)

	fmt.Println("Original slice", sicea)

	sliceCopy := make([]int, len(sicea))

	copy(sliceCopy, sicea)

	fmt.Println("Copied slice", sliceCopy)

	// var nilSlice []int

	for i, v := range sicea {
		fmt.Println(i, v)
	}

	fmt.Println("element at index 3", sliceCopy[3])

	// sliceCopy[3] = 200

	// fmt.Println("After modification", sliceCopy[3])

	fmt.Println("element at index 3", sicea[3])

	if slices.Equal(sicea, sliceCopy) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d at index %d of outer slice at index %d of inner slice\n", i+j, i, j)
		}
	}

	fmt.Println("2d: ", twoD)

	sicea = append(sicea, 500, 600, 700)

	fmt.Println("sicea", sicea)


	sliceNew := sicea[2:4]

	fmt.Println("sliceNew", sliceNew)

	fmt.Println("The capacity of sliceNew", cap(sliceNew))

	fmt.Println("The length of sliceNew", len(sliceNew))

	fmt.Println("The capacity of sicea", cap(sicea))


	
	

}
