package models

import (
	"encoding/json"
	"time"
)

type Purchase struct {
	Id          int        `json:"id"`
	Date        time.Time  `json:"date"`
	TotalPrice  float64    `json:"totalPrice"`
	PaymentDate *time.Time `json:"paymentDate"`
	CustomerId  int        `json:"customerId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Products    []struct {
		Id          int     `json:"id"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Quantity    int     `json:"quantity"`
	} `json:"products"`
}

func (purchase Purchase) MarshalJSON() ([]byte, error) {
	type Alias Purchase

	return json.Marshal(&struct {
		Date string `json:"date"`
		*Alias
	}{
		Date:  purchase.Date.Format("2006-01-02"),
		Alias: (*Alias)(&purchase),
	})
}
