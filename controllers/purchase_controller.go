package controllers

import (
	"context"
	"mercadinho/database"
	"mercadinho/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPurchases godoc
// @Summary     Get all registered purchases
// @Tags        Purchases
// @Produce     json
// @Success     200  {array}   models.Purchase
// @Failure     500  {object}  models.ErrorResponse
// @Router      /purchases [get]
func GetPurchases(ctx *gin.Context) {
	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, date, payment_date, customer_id, created_at, updated_at
		FROM purchases
	`)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	defer rows.Close()

	purchases := []models.Purchase{}

	for rows.Next() {
		var purchase models.Purchase

		err := rows.Scan(
			&purchase.ID,
			&purchase.Date,
			&purchase.PaymentDate,
			&purchase.CustomerID,
			&purchase.CreatedAt,
			&purchase.UpdatedAt,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		purchases = append(purchases, purchase)
	}

	ctx.IndentedJSON(http.StatusOK, purchases)
}

// GetPurchase godoc
// @Summary    Get a purchase by ID
// @Tags       Purchases
// @Produce    json
// @Success    200  {object}  models.Purchase
// @Failure    400  {object}  models.ErrorResponse
// @Failure    404  {object}  models.ErrorResponse
// @Router     /purchases/{id} [get]
func GetPurchase(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var purchase models.Purchase

	err = database.Pool.QueryRow(context.Background(), `
		SELECT id, date, payment_date, customer_id, created_at, updated_at
		FROM purchases
		WHERE id = $1
	`, id).Scan(
		&purchase.ID,
		&purchase.Date,
		&purchase.PaymentDate,
		&purchase.CustomerID,
		&purchase.CreatedAt,
		&purchase.UpdatedAt,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Purchase not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, purchase)
}

// TODO: Create routes to edit and remove purchase
