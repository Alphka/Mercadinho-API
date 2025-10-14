package dtos

type CreatePurchaseDTO struct {
	Date     string `json:"date" binding:"required" example:"2025-10-13"`
	Products []struct {
		Id       int `json:"id" binding:"required"`
		Quantity int `json:"quantity" binding:"required,gte=1"`
	} `json:"products" binding:"required,min=1"`
	PaymentDate *string `json:"paymentDate" example:"2025-10-13"`
}

type EditPurchaseDTO struct {
	Date string `json:"date" binding:"required" example:"2025-10-13"`
}
