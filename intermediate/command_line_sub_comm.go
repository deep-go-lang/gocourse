package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'firstsub' or 'secondsub' subcommands")
		os.Exit(1)
	}

	// Create subcommand flag sets
	subCommand1 := flag.NewFlagSet("firstsub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondsub", flag.ExitOnError)

	// Define flags for firstsub
	firstFlag := subCommand1.Bool("processing", false, "command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "byte lenght of result")
	userFlag1 := subCommand1.String("user", "guest", "username")

	// Define flags for secondsub
	flagc2 := subCommand2.String("language", "Go", "Enter the language")
	userFlag2 := subCommand2.String("user", "guest", "username")

	// Get the subcommand
	subcommand := os.Args[1]

	switch subcommand {
	case "firstsub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("username", *userFlag1)
		fmt.Println("firstsub command selected")
		fmt.Println("processing", *firstFlag)
		fmt.Println("bytes", *secondFlag)

	case "secondsub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("username", *userFlag2)
		fmt.Println("secondsub command selected")
		fmt.Println("language", *flagc2)

	default:
		fmt.Println("No subcommand selected")
		os.Exit(1)
	}
}
