package config

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Target struct {
	URL      string
	Interval time.Duration
}

func LoadTargets(filename string) ([]Target, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var targets []Target
	for i, record := range records {
		if len(record) < 2 {
			continue // Skip invalid lines
		}
		// Skip header if it exists and looks like a header
		if i == 0 && (record[0] == "url" || record[1] == "interval") {
			continue
		}

		url := record[0]
		intervalSec, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("invalid interval on line %d: %v", i+1, err)
		}

		targets = append(targets, Target{
			URL:      url,
			Interval: time.Duration(intervalSec) * time.Second,
		})
	}

	return targets, nil
}
