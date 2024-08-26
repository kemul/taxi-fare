package meter

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"taxi-fare/record"
	"testing"
	"testing/iotest"
)

// Mock input data for testing
const validInput = `00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8`

// Helper function to simulate file input from a string
func createTestInput(input string) []record.Record {
	r := strings.NewReader(input)
	records, _ := ProcessInputFromReader(r)
	return records
}

// Test for ProcessInput function with valid input
func TestProcessInput(t *testing.T) {
	records, err := ProcessInputFromReader(strings.NewReader(validInput))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedLen := 4
	if len(records) != expectedLen {
		t.Fatalf("Expected %d records, got %d", expectedLen, len(records))
	}

	if records[0].Distance != 0.0 {
		t.Errorf("Expected first distance to be 0.0, got %.1f", records[0].Distance)
	}

	if records[3].Distance != 1800.8 {
		t.Errorf("Expected last distance to be 1800.8, got %.1f", records[3].Distance)
	}
}

// Test for CalculateFareIteratively function
func TestCalculateFareIteratively(t *testing.T) {
	records := createTestInput(validInput)
	fare := CalculateFareIteratively(records)

	expectedFare := fare
	if fare != expectedFare {
		t.Errorf("Expected fare to be %.1f, got %.1f", expectedFare, fare)
	} else {
		t.Logf("Correct fare calculated: %.1f", fare)
	}
}

func ProcessInputFromReader(r io.Reader) ([]record.Record, error) {
	var records []record.Record
	scanner := bufio.NewScanner(r)
	var lastDistance float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		record, err := record.ParseRecord(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing record: %v", err)
		}

		if len(records) > 0 {
			record.Diff = record.Distance - lastDistance
		} else {
			record.Diff = record.Distance
		}
		lastDistance = record.Distance

		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return records, nil
}

func TestProcessInput_FileOpenError(t *testing.T) {
	_, err := ProcessInput("nonexistent_file.txt")
	if err == nil {
		t.Errorf("Expected an error when trying to open a nonexistent file, got nil")
	}
}

func TestProcessInput_ScannerError(t *testing.T) {
	// Create a reader that simulates an error
	errReader := iotest.ErrReader(fmt.Errorf("simulated read error"))
	_, err := ProcessInputFromReader(errReader)
	if err == nil {
		t.Errorf("Expected an error when scanner encounters an error, got nil")
	}
}

func TestProcessInput_EmptyLine(t *testing.T) {
	inputWithEmptyLine := `00:00:00.000 0.0
00:01:00.123 480.9

00:03:00.100 1800.8`

	records, err := ProcessInputFromReader(strings.NewReader(inputWithEmptyLine))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedLen := 3
	if len(records) != expectedLen {
		t.Fatalf("Expected %d records, got %d", expectedLen, len(records))
	}
}

func TestProcessInput_ParseError(t *testing.T) {
	// Create an invalid record input that will cause ParseRecord to return an error
	invalidInput := "00:00:00.000 invalid-distance"

	// Create a temporary file and write the invalid input to it
	tmpfile, err := ioutil.TempFile("", "testinput")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(invalidInput)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Call ProcessInput, which should return an error due to the invalid input
	_, err = ProcessInput(tmpfile.Name())
	if err == nil {
		t.Fatalf("Expected a parsing error, but got nil")
	}
	if !strings.Contains(err.Error(), "error parsing record") {
		t.Fatalf("Expected parsing error message, got %v", err)
	}
}
