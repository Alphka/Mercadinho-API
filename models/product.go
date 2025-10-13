package models

import "time"

type Product struct {
	ID          int        `json:"id"`
	Price       float64    `json:"preco"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
