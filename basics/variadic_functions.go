package basics

import "fmt"

func main() {

	statement, total := sum("The sum of the numbers is", 1, 2, 3, 4, 5)

	fmt.Println(statement, total)


	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sequence, total := addition(1, numbers...)

	fmt.Printf("for sequence %d, the total is %d\n", sequence, total)

	actresses := []string{"Aishwarya Rai", "Kareena Kapoor", "Deepika Padukone", "Priyanka Chopra"}

	celebrity("I love you", actresses...)

	

}

func sum(returnString string, nums... int) (string, int){
	total := 0
	for _, num := range nums {
		total += num
	}
	return returnString, total
}

func addition(sequence int, nums... int) (int, int){
	total := 0
	for _, num := range nums {
		total += num
	}
	return sequence, total
}

func celebrity(returnString string, actresses... string){
	for _, actress := range actresses {
		fmt.Println(returnString, actress)
	}
	
}
