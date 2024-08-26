package record

import (
	"testing"
	"time"
)

func TestParseRecord(t *testing.T) {
	// Test with valid input
	validLine := "12:34:56.789 1234.5"
	expectedTime, _ := time.Parse("15:04:05.000", "12:34:56.789")
	expectedDistance := 1234.5

	record, err := ParseRecord(validLine)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !record.Time.Equal(expectedTime) {
		t.Errorf("Expected time %v, got %v", expectedTime, record.Time)
	}

	if record.Distance != expectedDistance {
		t.Errorf("Expected distance %.1f, got %.1f", expectedDistance, record.Distance)
	}

	// Test with invalid time format
	invalidTimeLine := "invalid-time 1234.5"
	_, err = ParseRecord(invalidTimeLine)
	if err == nil {
		t.Errorf("Expected error for invalid time format, got nil")
	}

	// Test with invalid distance format
	invalidDistanceLine := "12:34:56.789 invalid-distance"
	_, err = ParseRecord(invalidDistanceLine)
	if err == nil {
		t.Errorf("Expected error for invalid distance format, got nil")
	}

	// Test with invalid input format
	invalidFormatLine := "12:34:56.789"
	_, err = ParseRecord(invalidFormatLine)
	if err == nil {
		t.Errorf("Expected error for invalid input format, got nil")
	}
}
