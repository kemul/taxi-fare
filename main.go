package main

import (
	"fmt"
	"log"
	"os"
	"taxi-fare/meter"
)

// Refactor core logic into a separate function for testability
func runTaxiFareCalculator(inputFilePath string) (int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	taxiMeter := meter.NewTaxiMeter()
	if err := taxiMeter.ProcessRecords(file); err != nil {
		return 0, fmt.Errorf("error processing records: %v", err)
	}

	fare := taxiMeter.CalculateFare()
	taxiMeter.DisplaySortedRecords()

	return int(fare), nil
}

func main() {
	fare, err := runTaxiFareCalculator("/root/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(fare)
}
