package meter

import (
	"strings"
	"taxi-fare/record"
	"testing"
	"time"
)

func TestNewTaxiMeter(t *testing.T) {
	taxiMeter := NewTaxiMeter()
	if taxiMeter == nil {
		t.Errorf("Expected a new TaxiMeter instance, got nil")
	}
}

func TestProcessRecords(t *testing.T) {
	taxiMeter := NewTaxiMeter()

	// Mock input data
	inputData := `00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8`

	inputFile := strings.NewReader(inputData)

	err := taxiMeter.ProcessRecords(inputFile)
	if err != nil {
		t.Errorf("Unexpected error in ProcessRecords: %v", err)
	}

	// Check if records were processed correctly
	if len(taxiMeter.records) != 4 {
		t.Errorf("Expected 4 records, got %d", len(taxiMeter.records))
	}

	// Verify the first record
	expectedFirstRecord := record.Record{
		Time:     time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		Distance: 0.0,
		Diff:     0.0,
	}

	if taxiMeter.records[0] != expectedFirstRecord {
		t.Errorf("Expected first record: %+v, got %+v", expectedFirstRecord, taxiMeter.records[0])
	}
}

func TestValidateFinalRecords(t *testing.T) {
	taxiMeter := NewTaxiMeter()

	// Mock input data
	inputData := `00:00:00.000 0.0
00:01:00.123 480.9`

	inputFile := strings.NewReader(inputData)

	err := taxiMeter.ProcessRecords(inputFile)
	if err != nil {
		t.Errorf("Unexpected error in ProcessRecords: %v", err)
	}

	// Validate final records
	err = taxiMeter.validateFinalRecords()
	if err != nil {
		t.Errorf("Unexpected error in validateFinalRecords: %v", err)
	}
}

func TestDisplaySortedRecords(t *testing.T) {
	taxiMeter := NewTaxiMeter()

	// Mock input data
	inputData := `00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8`

	inputFile := strings.NewReader(inputData)

	err := taxiMeter.ProcessRecords(inputFile)
	if err != nil {
		t.Errorf("Unexpected error in ProcessRecords: %v", err)
	}

	taxiMeter.DisplaySortedRecords()

	// Check if records are sorted by the Diff field
	if taxiMeter.records[0].Diff < taxiMeter.records[1].Diff {
		t.Errorf("Expected records to be sorted by Diff in descending order")
	}
}
