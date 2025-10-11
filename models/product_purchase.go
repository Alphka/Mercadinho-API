package models

type ProductPurchase struct {
	ID         int     `json:"id"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	PurchaseID int
	ProductID  int
}
