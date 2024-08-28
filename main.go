package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"taxi-fare/meter"
	"taxi-fare/record"

	"github.com/sirupsen/logrus"
)

// Initialize logrus logger
var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

var Run = run // Assign the actual function to the variable

func validateInput(input string) (string, error) {
	if input == "" {
		return "", errors.New("input file path cannot be empty")
	}
	return input, nil
}

func run(inputFilePath string) (float64, []record.Record, error) {
	// Input validation
	validatedPath, err := validateInput(inputFilePath)
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "input_validation",
			"path":  inputFilePath,
			"error": err.Error(),
		}).Error("Invalid input file path")
		return 0, nil, err
	}

	records, err := meter.ProcessInput(validatedPath)
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "file_processing",
			"path":  validatedPath,
			"error": err.Error(),
		}).Error("Failed to process input file")
		return 0, nil, err
	}

	fare := meter.CalculateFareIteratively(records)
	log.WithFields(logrus.Fields{
		"event": "fare_calculation",
		"fare":  fare,
	}).Info("Calculated fare successfully")

	return fare, records, nil
}

func MainLogic() int {
	fare, records, err := Run("input.txt")
	if err != nil {
		log.WithFields(logrus.Fields{
			"event": "main_logic",
			"error": err.Error(),
		}).Error("Error processing input in MainLogic")
		return 1
	}

	// Print the fare as an integer in regular format
	fmt.Printf("(iv) Output =============================\n")
	fmt.Printf("%d\n", int(fare))

	// Sort records by mileage difference in descending order
	sort.Slice(records, func(i, j int) bool {
		return records[i].Diff > records[j].Diff
	})

	// Print sorted records with distance and difference in regular format
	for _, rec := range records {
		fmt.Printf("%s %.1f %.1f\n",
			rec.Time.Format("15:04:05.000"),
			rec.Distance,
			rec.Diff,
		)
	}

	fmt.Printf("(iv) JSON Output =============================\n")
	log.WithFields(logrus.Fields{
		"event": "output",
		"fare":  int(fare),
	}).Info("(iv) Output Json")

	// Log sorted records with distance and difference in JSON format
	for _, rec := range records {
		log.WithFields(logrus.Fields{
			"time":     rec.Time.Format("15:04:05.000"),
			"distance": rec.Distance,
			"diff":     rec.Diff,
		}).Info("Processed record")
	}

	return 0
}

func main() {
	exitCode := MainLogic()
	os.Exit(exitCode)
}
