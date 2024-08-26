package main

import (
	"log"
	"os"
	"sort"
	"taxi-fare/meter"
	"taxi-fare/record"
)

var Run = run // Assign the actual function to the variable

func run(inputFilePath string) (float64, []record.Record, error) {
	records, err := meter.ProcessInput(inputFilePath)
	if err != nil {
		return 0, nil, err
	}

	fare := meter.CalculateFareIteratively(records)
	return fare, records, nil
}

func MainLogic() int {
	fare, records, err := Run("input.txt")
	if err != nil {
		log.Printf("Error processing input: %v", err)
		return 1
	}

	// Print the fare as an integer
	log.Printf("(iv) Output =============================\n")
	log.Printf("%d\n", int(fare))

	// Sort records by mileage difference in descending order
	sort.Slice(records, func(i, j int) bool {
		return records[i].Diff > records[j].Diff
	})

	// Print sorted records with distance and difference
	for _, rec := range records {
		log.Printf("%s %.1f %.1f\n",
			rec.Time.Format("15:04:05.000"),
			rec.Distance,
			rec.Diff,
		)
	}

	return 0
}

func main() {
	exitCode := MainLogic()
	os.Exit(exitCode)
}
