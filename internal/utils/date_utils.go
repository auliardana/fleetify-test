package utils

import (
	"errors"
	"time"
)

// ParseDateRange parses and validates start and end dates from strings.
// If a date string is empty, it defaults to the current date for startDate or the next day for endDate.
func ParseDateRange(startDateStr, endDateStr string) (time.Time, time.Time, error) {
	layout := "2006-01-02"

	// Default values
	if startDateStr == "" {
		startDateStr = time.Now().Truncate(24 * time.Hour).Format(layout)
	}
	if endDateStr == "" {
		endDateStr = time.Now().AddDate(0, 0, 1).Format(layout)
	}

	// Parse dates
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("invalid start_date format")
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("invalid end_date format")
	}

	if startDate.After(endDate) {
		return time.Time{}, time.Time{}, errors.New("start_date must be less than or equal to end_date")
	}

	return startDate, endDate, nil
}
