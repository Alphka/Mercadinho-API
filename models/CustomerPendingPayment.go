package models

import "time"

type CustomerPendingPayment struct {
	Id         int        `json:"id"`
	Date       time.Time  `json:"date"`
	TotalPrice float64    `json:"totalPrice"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
