package main

import (
	"log"
	"os"
	"taxi-fare/meter"
)

var Run = run // Assign the actual function to the variable

func run(inputFilePath string) (float64, error) {
	records, err := meter.ProcessInput(inputFilePath)
	if err != nil {
		return 0, err
	}

	fare := meter.CalculateFareIteratively(records)
	return fare, nil
}

func MainLogic() int {
	fare, err := Run("input.txt")
	if err != nil {
		log.Printf("Error processing input: %v", err)
		return 1
	}

	log.Printf("Total Fare: %d yen\n", int(fare))
	return 0
}

func main() {
	exitCode := MainLogic()
	os.Exit(exitCode)
}
