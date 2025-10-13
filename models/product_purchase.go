package models

import "time"

type ProductPurchase struct {
	Id         int        `json:"id"`
	Price      float64    `json:"price"`
	Quantity   int        `json:"quantity"`
	PurchaseId int        `json:"purchaseId"`
	ProductId  int        `json:"productId"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
