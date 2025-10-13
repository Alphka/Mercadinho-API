package helpers

import (
	"strings"
	"time"
)

func ParseDate(date string) *time.Time {
	date = strings.Trim(date, `"`)
	if date == "" {
		return nil
	}

	if t, err := time.Parse("2006-01-02", date); err == nil {
		return &t
	}

	if t, err := time.Parse(time.RFC3339, date); err == nil {
		return &t
	}

	// TODO: Handle error
	// fmt.Errorf("invalid date format: %s", date)
	return nil
}
