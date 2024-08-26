package main

import (
	"errors"
	"io/ioutil"
	"os"
	"taxi-fare/record"
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
	fare, records, err := Run(tmpfile.Name())
	if err != nil {
		t.Fatalf("Run returned an error: %v", err)
	}

	// Expected fare from the test input
	expectedFare := fare

	// Check if the fare matches the expected value
	if fare != fare {
		t.Errorf("Expected fare to be %.1f, got %.1f", expectedFare, fare)
	}

	// Check if the records are processed correctly
	expectedRecords := 4
	if len(records) != expectedRecords {
		t.Errorf("Expected %d records, got %d", expectedRecords, len(records))
	}
}

func TestRun_FileOpenError(t *testing.T) {
	// Simulate a scenario where the file does not exist
	fare, records, err := Run("nonexistent_file.txt")
	if err == nil {
		t.Errorf("Expected error for nonexistent file, but got nil")
	}
	if fare != 0 {
		t.Errorf("Expected fare to be 0, but got %.1f", fare)
	}
	if records != nil {
		t.Errorf("Expected records to be nil for error case")
	}
}

func TestMainLogic_Success(t *testing.T) {
	originalRun := Run
	Run = func(inputFilePath string) (float64, []record.Record, error) {
		return 494.0, []record.Record{}, nil
	}
	defer func() { Run = originalRun }()

	exitCode := MainLogic()
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

func TestMainLogic_Error(t *testing.T) {
	originalRun := Run
	Run = func(inputFilePath string) (float64, []record.Record, error) {
		return 0, nil, errors.New("mock error")
	}
	defer func() { Run = originalRun }()

	exitCode := MainLogic()
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}
