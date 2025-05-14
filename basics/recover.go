package basics


import "fmt"

func main() {
    process()
    fmt.Println("program recovered")

}

func process() {
    defer func() {
        r := recover()
        if r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fmt.Println("Process started")
    panic("Something went wrong")
    fmt.Println("This will not be printed")
}
