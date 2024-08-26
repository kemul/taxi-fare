package meter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"taxi-fare/record"
)

const (
	baseFare     = 400.0
	farePer400m  = 40.0
	baseDistance = 1000.0
)

func ProcessInput(filePath string) ([]record.Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	var records []record.Record
	scanner := bufio.NewScanner(file)
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

func CalculateFareIteratively(records []record.Record) float64 {
	fare := baseFare
	lastDistance := 0.0

	for i, record := range records {
		log.Printf("=========================================================\n")
		log.Printf("Processing line: %v", record) // Log each line as it is processed
		if i == 0 {
			log.Printf("Step %d: Initial fare: %d yen for up to 1 km.\n", i+1, int(fare))
		} else {
			log.Printf("Step %d: Current Distance: %.1f meters\n", i+1, record.Distance)
			if record.Distance > baseDistance {
				extraDistance := record.Distance - baseDistance

				// Only consider the distance beyond the last recorded distance
				if lastDistance > baseDistance {
					extraDistance = record.Distance - lastDistance
				}

				// Calculate additional fare
				numUnits := extraDistance / 400.0
				additionalFare := numUnits * farePer400m
				fare += additionalFare

				log.Printf("Additional distance beyond 1 km: %.1f meters\n", extraDistance)
				log.Printf("Number of 400m units: %.2f\n", numUnits)
				log.Printf("Additional fare: %.2f yen (%.2f * %.2f)\n", additionalFare, numUnits, farePer400m)
				log.Printf("Total fare after this step: %d yen\n", int(fare))
			} else {
				log.Printf("Still within the first 1 km, no additional fare. Fare remains: %d yen\n", int(fare))
			}
		}

		lastDistance = record.Distance
	}

	return fare
}
