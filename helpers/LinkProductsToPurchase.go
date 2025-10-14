package helpers

import (
	"context"
	"fmt"
	"mercadinho/database"
	"mercadinho/models"
)

func LinkProductsToPurchase(purchase *models.Purchase, productIds *[]int) error {
	var productPurchases []struct {
		Id       int     `json:"id"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	if productIds == nil {
		rows, err := database.Pool.Query(context.Background(), `
			SELECT product_id, price, quantity
			FROM product_purchase
			WHERE purchase_id = $1
		`, purchase.Id)
		if err != nil {
			return fmt.Errorf("failed to get purchase products")
		}
		defer rows.Close()

		for rows.Next() {
			var productPurchase struct {
				Id       int     `json:"id"`
				Price    float64 `json:"price"`
				Quantity int     `json:"quantity"`
			}

			err := rows.Scan(
				&productPurchase.Id,
				&productPurchase.Price,
				&productPurchase.Quantity,
			)
			if err != nil {
				return err
			}

			productPurchases = append(productPurchases, productPurchase)
		}

		(*purchase).Products = productPurchases

		return nil
	}

	rows, err := database.Pool.Query(context.Background(), `
		SELECT product_id, price, quantity
		FROM product_purchase
		WHERE purchase_id = $1 AND product_id = ANY($2)
	`, purchase.Id, *productIds)
	if err != nil {
		return fmt.Errorf("failed to get purchase products")
	}
	defer rows.Close()

	for rows.Next() {
		var productPurchase struct {
			Id       int     `json:"id"`
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		}

		err := rows.Scan(
			&productPurchase.Id,
			&productPurchase.Price,
			&productPurchase.Quantity,
		)
		if err != nil {
			return err
		}

		productPurchases = append(productPurchases, productPurchase)
	}

	(*purchase).Products = productPurchases

	return nil
}
