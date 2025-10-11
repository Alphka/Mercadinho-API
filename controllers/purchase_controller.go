package controllers

import (
	"mercadinho/models"
	"time"
)

var Purchases = []models.Purchase{
	{
		ID:          1,
		Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TotalPrice:  100.00,
		PaymentDate: nil,
		CustomerID:  1,
	},
}
