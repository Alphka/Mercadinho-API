package dtos

type CreatePurchaseDTO struct {
	Date        string  `json:"date" binding:"required" example:"2025-10-13"`
	Products    []int   `json:"products" binding:"required,min=1" example:"[1]"`
	PaymentDate *string `json:"paymentDate" example:"2025-10-13"`
}

type EditPurchaseDTO struct {
	Date string `json:"date" binding:"required" example:"2025-10-13"`
}
