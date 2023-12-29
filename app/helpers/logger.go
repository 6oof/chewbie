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
func initLogger() {
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

// LogInfo writes an info-level log message.
func LogInfo(message string) {
	mu.Lock()
	defer mu.Unlock()
	once.Do(initLogger)
	logger.Println("INFO:", message)
}

// LogError writes an error-level log message.
func LogError(err error, message string) {
	mu.Lock()
	defer mu.Unlock()
	once.Do(initLogger)
	logger.Println("ERROR:", message, "-", err)
}
