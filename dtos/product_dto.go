package dtos

type CreateProductDTO struct {
	Price       float64 `json:"price" example:"10.50"`
	Description string  `json:"description"`
}

type EditProductDTO struct {
	Price       float64 `json:"price" example:"10.50"`
	Description string  `json:"description"`
}
