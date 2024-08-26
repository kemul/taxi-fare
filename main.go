package main

import (
	"log"
	"taxi-fare/meter"
)

func main() {
	records, err := meter.ProcessInput("input.txt")
	if err != nil {
		log.Fatalf("Error processing input: %v", err)
	}

	fare := meter.CalculateFareIteratively(records)
	log.Printf("Total Fare: %d yen\n", int(fare))
}
