package basics


import "fmt"

func main() {

    // var sum int = addition(10, 20)

    fmt.Println(addition(100, 500))

    crush := func() {
        fmt.Println("Gorgeous Kriti Sanon!")
    }

    crush()

    summation := addition
    
    value := summation(40, 10)

    fmt.Println("The sum of 40 and 10 is", value)

    love := loveStatement("Tapsee")

    func() {
        fmt.Println("I love you", love)
    }()

    myResult := applyOperation(10, 20, addition)

    fmt.Println("The result of 10 and 20 is", myResult)

    multiplier := createMultiplier(5)

    result := multiplier(10)

    fmt.Println("The result of 10 and 5 is", result)    

}

func loveStatement(name string) string{
    return name
}

func addition(x int, y int) int {
    return x + y
}

func applyOperation(x int, y int, operation func(int, int) int) int {
    return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}


