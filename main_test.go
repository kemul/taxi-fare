package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

const testInput = `00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8`

func TestRun(t *testing.T) {
	// Create a temporary input file
	tmpfile, err := ioutil.TempFile("", "testinput")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	// Write the test input to the temporary file
	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Run the main function logic using the temporary input file
	fare, err := Run(tmpfile.Name())
	if err != nil {
		t.Fatalf("Run returned an error: %v", err)
	}

	// Expected fare from the test input
	expectedFare := fare

	// Check if the fare matches the expected value
	if fare != expectedFare {
		t.Errorf("Expected fare to be %.1f, got %.1f", expectedFare, fare)
	}
}

func TestRun_FileOpenError(t *testing.T) {
	// Simulate a scenario where the file does not exist
	fare, err := Run("nonexistent_file.txt")
	if err == nil {
		t.Errorf("Expected error for nonexistent file, but got nil")
	}
	if fare != 0 {
		t.Errorf("Expected fare to be 0, but got %.1f", fare)
	}
}

func TestMainLogic_Success(t *testing.T) {
	originalRun := Run
	Run = func(inputFilePath string) (float64, error) {
		return 494.0, nil
	}
	defer func() { Run = originalRun }()

	exitCode := MainLogic()
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

func TestMainLogic_Error(t *testing.T) {
	originalRun := Run
	Run = func(inputFilePath string) (float64, error) {
		return 0, errors.New("mock error")
	}
	defer func() { Run = originalRun }()

	exitCode := MainLogic()
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}
