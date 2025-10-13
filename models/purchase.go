package models

import "time"

type Purchase struct {
	ID          int        `json:"id"`
	Date        time.Time  `json:"date"`
	TotalPrice  float64    `json:"totalPrice"`
	PaymentDate *time.Time `json:"paymentDate"`
	CustomerID  int        `json:"customerId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
