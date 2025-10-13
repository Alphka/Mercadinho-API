package models

import "time"

type ProductPurchase struct {
	ID         int        `json:"id"`
	Price      float64    `json:"price"`
	Quantity   int        `json:"quantity"`
	PurchaseID int        `json:"purchaseId"`
	ProductID  int        `json:"productId"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
