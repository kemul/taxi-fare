package record

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestParseRecord(t *testing.T) {
	tests := []struct {
		input          string
		expectedRecord Record
		expectError    bool
	}{
		{
			input: "00:01:00.123 480.9",
			expectedRecord: Record{
				Time:     time.Date(0, 1, 1, 0, 1, 0, 123000000, time.UTC),
				Distance: 480.9,
				Diff:     0, // Diff is not set by ParseRecord
			},
			expectError: false,
		},
		{
			input:       "00:01:00.123",
			expectError: true,
		},
		{
			input:       "invalid time 480.9",
			expectError: true,
		},
		{
			input:       "00:01:00.123 invalid",
			expectError: true,
		},
	}

	for _, test := range tests {
		record, err := ParseRecord(test.input)

		if test.expectError && err == nil {
			t.Errorf("Expected an error for input %q, but got none", test.input)
		}

		if !test.expectError {
			if err != nil {
				t.Errorf("Did not expect an error for input %q, but got %v", test.input, err)
			}

			if record.Time != test.expectedRecord.Time {
				t.Errorf("Expected time %v, got %v", test.expectedRecord.Time, record.Time)
			}

			if record.Distance != test.expectedRecord.Distance {
				t.Errorf("Expected distance %v, got %v", test.expectedRecord.Distance, record.Distance)
			}
		}
	}
}

func TestPrintRecord(t *testing.T) {
	record := Record{
		Time:     time.Date(0, 1, 1, 0, 1, 0, 123000000, time.UTC),
		Distance: 480.9,
		Diff:     100,
	}

	expectedOutput := `{"time":"0000-01-01T00:01:00.123Z","distance":480.9,"diff":100}`

	// Capture the output of PrintRecord
	output := captureOutput(func() {
		record.PrintRecord()
	})

	if output != expectedOutput+"\n" { // Adding \n because fmt.Println adds a newline
		t.Errorf("Expected output %q, but got %q", expectedOutput, output)
	}
}

// Helper function to capture output from PrintRecord
func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
