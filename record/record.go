package record

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Time     time.Time
	Distance float64
	Diff     float64
}

func ParseRecord(line string) (Record, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Record{}, fmt.Errorf("invalid input format")
	}

	parsedTime, err := time.Parse("15:04:05.000", parts[0])
	if err != nil {
		return Record{}, fmt.Errorf("invalid time format")
	}

	distance, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Record{}, fmt.Errorf("invalid distance format")
	}

	return Record{Time: parsedTime, Distance: distance}, nil
}
