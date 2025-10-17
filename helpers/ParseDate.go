package helpers

import (
	"errors"
	"strings"
	"time"
)

var DateLayouts = []string{
	time.RFC3339,
	"2006-01-02",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	time.RFC1123,
	time.RFC1123Z,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.RFC3339Nano,
}

func ParseDate(date string) (*time.Time, error) {
	date = strings.TrimSpace(strings.Trim(date, `"`))
	if date == "" {
		return nil, errors.New("invalid date")
	}

	for _, layout := range DateLayouts {
		if t, err := time.Parse(layout, date); err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("invalid date format")
}
