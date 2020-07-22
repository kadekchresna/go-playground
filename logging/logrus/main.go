package main

import (
	"fmt"
	"os"

	"github.com/kadekchresna/playground/logging/other"
	log "github.com/sirupsen/logrus"
)

func main() {

	// open a file
	f, err := os.OpenFile("testlogrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// log.WithFields(log.Fields{
	// 	"Animal": "Logrus",
	// }).Info("A logrus appears")
	other.Other()
}
