package utils

import (
	"log"
	"os"
)

func SetupLogger() {
	// Create or open a log file (it appends to the file if it already exists)
	logFile, err := os.OpenFile("data/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(logFile)

	// Optional: Log the date and time in each log entry
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
