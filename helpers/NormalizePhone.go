package helpers

import (
	"errors"
	"regexp"
	"strings"
)

func NormalizePhone(phone string) (string, error) {
	regex := regexp.MustCompile(`\D`)
	cleaned := regex.ReplaceAllString(strings.TrimSpace(phone), "")

	if cleaned == "" {
		return "", errors.New("invalid phone number")
	}

	length := len(cleaned)

	switch {
	case length == 13:
		return cleaned, nil
	case length >= 8 && length <= 11:
		return "55" + cleaned, nil
	case length == 12:
		return "", errors.New("invalid phone number length (12 digits is not allowed)")
	default:
		return "", errors.New("invalid phone number format")
	}
}
