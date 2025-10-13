package models

import "time"

type Product struct {
	Id          int        `json:"id"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
