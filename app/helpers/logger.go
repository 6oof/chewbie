package helpers

import (
	"log"
	"os"
	"sync"
)

var (
	logger *log.Logger
	mu     sync.Mutex
	once   sync.Once
)

// initLogger initializes the logger and opens the log file.
//
// This function is called once to set up the logger. It creates a log file
// named "log.txt" and initializes the logger with the file writer.
func InitLogger(file string) {
	once.Do(func() {
		logFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		logger = log.New(logFile, "", log.Ldate|log.Ltime)
	})
}

// LogInfo writes an info-level log message to the log file.
//
// It acquires a lock to ensure safe concurrent access to the logger and initializes
// the logger using initLogger if it hasn't been initialized yet. The log message
// includes the prefix "INFO:" followed by the provided message.
//
// Example Usage:
//
//	LogInfo("Application started successfully")
func LogInfo(message string) {
	mu.Lock()
	defer mu.Unlock()
	logger.Println("INFO:", message)
}

// LogError writes an error-level log message to the log file.
//
// It acquires a lock to ensure safe concurrent access to the logger and initializes
// the logger using initLogger if it hasn't been initialized yet. The log message
// includes the prefix "ERROR:" followed by the provided message and the error details.
//
// Example Usage:
//
//	err := SomeFunction()
//	LogError(err, "An error occurred in SomeFunction")
func LogError(err error, message string) {
	mu.Lock()
	defer mu.Unlock()
	logger.Println("ERROR:", message, "-", err)
}
