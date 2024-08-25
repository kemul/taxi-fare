package main

import (
	"fmt"
	"log"
	"os"
	"taxi-fare/meter"
)

func main() {
	taxiMeter := meter.NewTaxiMeter()
	if err := taxiMeter.ProcessRecords(os.Stdin); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fare := taxiMeter.CalculateFare()
	fmt.Println(int(fare))

	taxiMeter.DisplaySortedRecords()
}
