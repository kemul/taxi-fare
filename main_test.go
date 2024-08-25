package main

import (
	"os"
	"testing"
)

func TestRunTaxiFareCalculator(t *testing.T) {
	// Mock input data
	inputData := `00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8`

	// Create a temporary file with the mock input data
	tmpfile, err := os.CreateTemp("", "input.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after the test

	if _, err := tmpfile.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Run the function with the temporary file as input
	fare, err := runTaxiFareCalculator(tmpfile.Name())
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Validate the calculated fare
	expectedFare := 480
	tolerance := 1 // Since fare is an int, we might use an int tolerance
	if abs(fare-expectedFare) > tolerance {
		t.Errorf("Expected fare: %d, got: %d", expectedFare, fare)
	}
}

// Helper function to calculate absolute value of an int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestRunTaxiFareCalculatorWithInvalidFile(t *testing.T) {
	// Pass an invalid file path
	_, err := runTaxiFareCalculator("/invalid/path/input.txt")
	if err == nil {
		t.Errorf("Expected an error for an invalid file path, but got none")
	}
}

func TestRunTaxiFareCalculatorWithInvalidData(t *testing.T) {
	// Mock invalid input data
	inputData := `invalid data`

	// Create a temporary file with the mock invalid data
	tmpfile, err := os.CreateTemp("", "input.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after the test

	if _, err := tmpfile.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Run the function with the temporary file as input
	_, err = runTaxiFareCalculator(tmpfile.Name())
	if err == nil {
		t.Errorf("Expected an error for invalid input data, but got none")
	}
}
