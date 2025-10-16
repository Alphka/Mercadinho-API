package models

import "time"

type Customer struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	Address      string     `json:"address"`
	Number       string     `json:"number"`
	Complement   *string    `json:"complement"`
	Neighborhood string     `json:"neighborhood"`
	City         string     `json:"city"`
	State        string     `json:"state"`
	ZIPCode      string     `json:"zipCode"`
	PaymentDay   int        `json:"paymentDay"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}
