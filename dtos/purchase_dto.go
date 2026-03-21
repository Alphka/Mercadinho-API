package dtos

import "strings"

type CreatePurchaseDTO struct {
	Date     string `json:"date" binding:"required,date" example:"2025-10-13"`
	Products []struct {
		Id       int `json:"id" binding:"required"`
		Quantity int `json:"quantity" binding:"required,gte=1"`
	} `json:"products" binding:"required,min=1"`
	PaymentDate *string `json:"paymentDate" binding:"omitempty,date" example:"2025-10-13"`
}

func (dto *CreatePurchaseDTO) Sanitize() {
	dto.Date = strings.TrimSpace(dto.Date)
}

type EditPurchaseDTO struct {
	Date string `json:"date" binding:"required,date" example:"2025-10-13"`
}

func (dto *EditPurchaseDTO) Sanitize() {
	dto.Date = strings.TrimSpace(dto.Date)
}
