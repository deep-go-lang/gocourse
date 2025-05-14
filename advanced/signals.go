package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)

	// Create a channel to receive signals
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Register for signals that can be caught in Windows
	signal.Notify(sigs, 
		syscall.SIGINT,    // Ctrl+C
		syscall.SIGTERM,   // Termination request
		syscall.SIGHUP,
	)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	go func() {

		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal...")
				os.Exit(0)
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}

		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT signal")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Received SIGTERM signal")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP signal")
		// 	}
		// 	fmt.Println("Graceful shutdown initiated...")
		// 	os.Exit(0)
		// }
	}()

	// fmt.Println("Waiting for signal...")
	// Block main goroutine
	select {}
}
