package utils

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
	"time"
)

func TestLogError(t *testing.T) {
	// Capture the log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil) // Restore the original output

	// Define an error to log
	err := errors.New("test error")

	// Call the function
	LogError(err)

	// Check the log output
	logOutput := buf.String()

	// Check that the log output contains the correct error message
	if !strings.Contains(logOutput, `"error":"test error"`) {
		t.Errorf("Expected log output to contain error message, got: %s", logOutput)
	}

	// Check that the log output contains the current time (within a tolerance)
	currentTime := time.Now().Format(time.RFC3339)
	if !strings.Contains(logOutput, `"time":"`+currentTime[:19]) { // Match the date and hour, ignore minutes and seconds
		t.Errorf("Expected log output to contain current time, got: %s", logOutput)
	}
}
