package record

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	LevelName string    `json:"level_name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func FromJSON(data string) (Record, error) {
	result := Record{}

	err := json.Unmarshal([]byte(data), &result)

	return result, err
}

func FromCSV(data string) (Record, error) {
	result := Record{}

	cols := strings.Split(data, ",")
	if len(cols) != 3 {
		return result, fmt.Errorf("expected CSV row with 3 columns, got %v", len(cols))
	}

	// parse timestamp
	ts, err := time.Parse(time.RFC3339, cols[0])
	if err != nil {
		return result, fmt.Errorf("could not parse Timestamp from CSV row %s", err.Error())
	}
	result.Timestamp = ts

	// parse level name
	result.LevelName = cols[1]

	// parse Value
	i, err := strconv.Atoi(cols[2])
	if err != nil {
		return result, fmt.Errorf("could not parse Value from CSV row %s", err.Error())
	}
	result.Value = i

	return result, err
}
