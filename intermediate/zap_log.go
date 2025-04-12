package intermediate

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)



func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Hello, Kriti!")

	fmt.Println("--------------------------------")

	logger.Info("User logged in", zap.String("username", "Kriti"), zap.String("method", "GET"))

}
