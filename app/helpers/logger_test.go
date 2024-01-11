package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestInitLogger(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Set log file path to the temporary directory
	logFilePath := tempDir + "/log.txt"

	// Initialize the logger
	InitLogger(logFilePath)

	// Ensure logger initialization works as expected
	assertFileExists(t, logFilePath, "Expected log file to be created")

}

func TestLogInfo(t *testing.T) {
	// Capture the logger output
	var buf bytes.Buffer
	logger = nil

	// Redirect logger output to buffer
	logger = log.New(&buf, "", log.Ldate|log.Ltime)

	// Log an info message
	LogInfo("Test Info Message")

	// Check if the log contains the expected content
	expectedContent := "INFO: Test Info Message"
	actualContent := buf.String()
	assertStringContains(t, actualContent, expectedContent, "LogInfo did not produce the expected log entry")
}

func TestLogError(t *testing.T) {
	// Capture the logger output
	var buf bytes.Buffer
	logger = nil

	// Redirect logger output to buffer
	logger = log.New(&buf, "", log.Ldate|log.Ltime)

	// Log an error message
	err := errors.New("Test Error")
	LogError(err, "Test Error Message")

	// Check if the log contains the expected content
	expectedContent := "ERROR: Test Error Message - Test Error"
	actualContent := buf.String()
	assertStringContains(t, actualContent, expectedContent, "LogError did not produce the expected log entry")
}

func TestConcurrentLogging(t *testing.T) {
	// Capture the logger output
	var buf bytes.Buffer
	logger = nil

	// Redirect logger output to buffer
	logger = log.New(&buf, "", log.Ldate|log.Ltime)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines to simulate concurrent access
	numGoroutines := 100

	// Simulate concurrent log entries
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			LogInfo(fmt.Sprintf("Concurrent Message %d", index))
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if the number of log entries is equal to the expected count
	expectedLogEntries := numGoroutines
	actualLogEntries := strings.Count(buf.String(), "INFO:")
	if actualLogEntries != expectedLogEntries {
		t.Errorf("Expected %d log entries, but got %d", expectedLogEntries, actualLogEntries)
	}
}

// Utility functions for testing

func assertFileExists(t *testing.T, path string, message string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("%s: %s does not exist", message, path)
	}
}

func assertFileNotExists(t *testing.T, path string, message string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("%s: %s should not exist", message, path)
	}
}

func assertStringContains(t *testing.T, actual, expected, message string) {
	t.Helper()
	if !strings.Contains(actual, expected) {
		t.Fatalf("%s: expected content '%s' not found in actual string '%s'", message, expected, actual)
	}
}
