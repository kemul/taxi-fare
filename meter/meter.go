package meter

import (
	"bufio"
	"errors"
	"io"
	"math"
	"sort"
	"taxi-fare/record"
	"taxi-fare/utils"
	"time"
)

const (
	baseFare        = 400.0
	farePer400m     = 40.0
	farePer350m     = 40.0
	baseDistance    = 1000.0
	midDistance     = 10000.0
	maxTimeInterval = 5 * time.Minute
)

type TaxiMeter struct {
	records []record.Record
}

func NewTaxiMeter() *TaxiMeter {
	return &TaxiMeter{}
}

func (tm *TaxiMeter) ProcessRecords(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // Skip blank lines instead of returning an error
		}

		rec, err := record.ParseRecord(line)
		if err != nil {
			utils.LogError(err) // Log the error but continue processing
			continue
		}

		if len(tm.records) > 0 {
			if err := tm.validateRecord(rec); err != nil {
				utils.LogError(err) // Log the error but continue processing
				continue
			}
			rec.Diff = rec.Distance - tm.getLastRecord().Distance
		} else {
			rec.Diff = rec.Distance
		}

		tm.records = append(tm.records, rec)
	}

	return tm.validateFinalRecords()
}

func (tm *TaxiMeter) validateRecord(rec record.Record) error {
	lastRecord := tm.getLastRecord()
	timeDiff := rec.Time.Sub(lastRecord.Time)

	if timeDiff < 0 {
		return errors.New("past time detected")
	}

	if timeDiff > maxTimeInterval {
		return errors.New("time interval greater than 5 minutes")
	}

	if rec.Distance < lastRecord.Distance {
		return errors.New("distance cannot decrease")
	}

	return nil
}

func (tm *TaxiMeter) validateFinalRecords() error {
	if len(tm.records) < 2 {
		return errors.New("insufficient data: less than two records")
	}

	if tm.getLastRecord().Distance == 0.0 {
		return errors.New("total distance is zero")
	}

	return nil
}

func (tm *TaxiMeter) CalculateFare() float64 {
	totalDistance := tm.getLastRecord().Distance
	fare := baseFare

	switch {
	case totalDistance > midDistance:
		fare += calculateFareSegment(midDistance-baseDistance, 400, farePer400m)
		fare += calculateFareSegment(totalDistance-midDistance, 350, farePer350m)
	case totalDistance > baseDistance:
		fare += calculateFareSegment(totalDistance-baseDistance, 400, farePer400m)
	}

	return math.Round(fare*100) / 100 // Round the final fare to two decimal places
}

func calculateFareSegment(distance, unit, rate float64) float64 {
	return math.Round((distance/unit)*rate*100) / 100
}

func (tm *TaxiMeter) DisplaySortedRecords() {
	sort.Slice(tm.records, func(i, j int) bool {
		return tm.records[i].Diff > tm.records[j].Diff
	})

	for _, rec := range tm.records {
		rec.PrintRecord()
	}
}

func (tm *TaxiMeter) getLastRecord() record.Record {
	return tm.records[len(tm.records)-1]
}
