package intermediate

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("Hello, World!")

	log.SetPrefix("ERROR: ")
	log.Println("This is an error message")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("This is a log message")

	// log.Fatal("This is a fatal message")

	// log.Panic("This is a panic message")

	fmt.Println("--------------------------------")

	infoLogger.Println("This is an info message")
	errorLogger.Println("This is an error message")
	warnLogger.Println("This is a warning message")

	fmt.Println("--------------------------------")

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	multiLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	var (
		infoLogger1  = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger1 = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		warnLogger1  = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	)

	multiLogger.Println("This is a new log message")

	infoLogger1.Println("This is an info message")
	errorLogger1.Println("This is an error message")
	warnLogger1.Println("This is a warning message")

	fmt.Println("--------------------------------")
	
	
	
	
	

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
)
