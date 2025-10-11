package models

type Product struct {
	ID          int     `json:"id"`
	Price       float64 `json:"preco"`
	Description string  `json:"description"`
}
