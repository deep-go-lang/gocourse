package intermediate

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message")
	log.Error("This is an error message")
	log.Warn("This is a warning message")

	log.WithFields(logrus.Fields{
		"username": "Urvashi Rautela",
		"method": "GET",}).Info("User logged in")

	
}