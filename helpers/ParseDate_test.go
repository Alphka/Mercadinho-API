package helpers

import "testing"

func TestParseDate(t *testing.T) {
	valid := []string{
		"2026-10-24",
		"2026-10-24T15:04:05",
		"2026-10-24 15:04:05",
		"2026-10-24T15:04:05Z",
		`"2026-10-24"`,
		"  2026-10-24  ",
	}

	for _, date := range valid {
		t.Run("accepts "+date, func(t *testing.T) {
			if _, err := ParseDate(date); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}

	invalid := []string{"", "   ", "not a date", "2026-13-45"}

	for _, date := range invalid {
		t.Run("rejects "+date, func(t *testing.T) {
			if _, err := ParseDate(date); err == nil {
				t.Error("expected an error")
			}
		})
	}
}
