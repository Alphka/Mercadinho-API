package dtos

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type customDate struct {
	time.Time
}

func (customDate *customDate) UnmarshalJSON(b []byte) error {
	string := strings.Trim(string(b), `"`)
	if string == "" {
		return nil
	}

	if time, err := time.Parse("2006-01-02", string); err == nil {
		customDate.Time = time
		return nil
	}

	if time, err := time.Parse(time.RFC3339, string); err == nil {
		customDate.Time = time
		return nil
	}

	return fmt.Errorf("invalid date format: %s", string)
}

func (customDate customDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(customDate.Time.Format("2006-01-02"))
}

type CreatePurchaseDTO struct {
	Date        customDate  `json:"date" binding:"required"`
	Products    []int       `json:"products" binding:"required,min=1"`
	PaymentDate *customDate `json:"paymentDate"`
}

type EditPurchaseDTO struct {
	Date time.Time `json:"date" binding:"required"`
}
