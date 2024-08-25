package record

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const timeLayout = "15:04:05.000"

type Record struct {
	Time     time.Time `json:"time"`
	Distance float64   `json:"distance"`
	Diff     float64   `json:"diff"`
}

func ParseRecord(line string) (Record, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return Record{}, errors.New("invalid format")
	}

	parsedTime, err := time.Parse(timeLayout, parts[0])
	if err != nil {
		return Record{}, errors.New("invalid time format")
	}

	distance, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return Record{}, errors.New("invalid distance format")
	}

	return Record{Time: parsedTime, Distance: distance}, nil
}

func (r *Record) PrintRecord() {
	recordJSON, _ := json.Marshal(r)
	fmt.Println(string(recordJSON))
}
