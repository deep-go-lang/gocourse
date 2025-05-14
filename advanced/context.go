package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

// func main() {

// 	todoContext := context.TODO()

// 	contextBackground := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "Urvashi")

// 	fmt.Println(ctx)

// 	fmt.Println(ctx.Value("name"))

// 	fmt.Println("--------------------------------")

// 	ctx = context.WithValue(contextBackground, "name", "Aiswharya")

// 	fmt.Println(ctx)

// 	fmt.Println(ctx.Value("name"))

// }

// func evenOrOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Context cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("Even number: %d", num)
// 		}
// 		return fmt.Sprintf("Odd number: %d", num)
// 	}
// }

// func demonstrateDefer() {
// 	fmt.Println("\nDemonstrating defer:")

// 	// Without defer - might forget to cancel
// 	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*1)
// 	fmt.Println("First context result:", evenOrOdd(ctx1, 5))
// 	cancel1() // Manual cancellation - easy to forget!

// 	// With defer - guaranteed cleanup
// 	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*1)
// 	defer cancel2() // This will be called when demonstrateDefer returns
// 	fmt.Println("Second context result:", evenOrOdd(ctx2, 6))

// 	fmt.Println("Contexts created")
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("Function ending - cancel2 will be called automatically")
// }

// func main() {
// 	ctx := context.TODO()

// 	result := evenOrOdd(ctx, 9)
// 	fmt.Println(result)

// 	fmt.Println("--------------------------------")

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
// 	defer cancel()

// 	result = evenOrOdd(ctx, 10)
// 	fmt.Println(result)

// 	fmt.Println("Starting 2 second sleep...")
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("Sleep finished, context should be cancelled now")

// 	result = evenOrOdd(ctx, 11)
// 	fmt.Println(result)

// 	demonstrateDefer()
// }

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, exiting", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestID", "1234567890")
	ctx = context.WithValue(ctx, "name", "Urvashi")
	ctx = context.WithValue(ctx, "ip", "190.155.1.1")
	ctx = context.WithValue(ctx, "os", "Windows")

	go doWork(ctx)

	time.Sleep(time.Second * 3)

	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("Request ID not found")
	}

	logWithContext(ctx, "Hello, World!")
	
}


func logWithContext(ctx context.Context, msg string) {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}
	log.Printf("requestID: %s - %s\n", requestID, msg)
}
