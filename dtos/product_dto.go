package dtos

import "strings"

type CreateProductDTO struct {
	Price       float64 `json:"price" binding:"price" example:"10.50"`
	Description string  `json:"description" binding:"required,min=3,max=255"`
}

func (dto *CreateProductDTO) Sanitize() {
	dto.Description = strings.TrimSpace(dto.Description)
}

type EditProductDTO struct {
	Price       float64 `json:"price" binding:"required,price" example:"10.50"`
	Description string  `json:"description" binding:"required,min=3,max=255"`
}

func (dto *EditProductDTO) Sanitize() {
	dto.Description = strings.TrimSpace(dto.Description)
}
