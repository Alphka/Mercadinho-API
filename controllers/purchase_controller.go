package controllers

import (
	"context"
	"mercadinho/database"
	"mercadinho/dtos"
	"mercadinho/helpers"
	"mercadinho/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPurchases godoc
// @Summary     Get all registered purchases
// @Tags        Purchases
// @Produce     json
// @Param       page           query  int  false  "Page number" default(1)
// @Param       pageSize       query  int  false  "Number of items per page" default(20)
// @Success     200  {array}   models.PaginatedResponse[models.Purchase]
// @Failure     500  {object}  models.ErrorResponse
// @Router      /purchases [get]
func GetPurchases(ctx *gin.Context) {
	var totalItems int

	err := database.Pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM purchases`).Scan(&totalItems)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	page, pageSize, totalPages := helpers.ValidatePagination(ctx, totalItems)

	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, date, payment_date, customer_id, created_at, updated_at
		FROM purchases
		ORDER BY created_at DESC
		OFFSET (($1 - 1) * $2)
		LIMIT $2
	`, page, pageSize)
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
			&purchase.Id,
			&purchase.Date,
			&purchase.PaymentDate,
			&purchase.CustomerId,
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

	for index := range purchases {
		if err := helpers.LinkProductsToPurchase(&purchases[index], nil); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
	}

	ctx.IndentedJSON(http.StatusOK, models.PaginatedResponse[models.Purchase]{
		Data:       purchases,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	})
}

// GetPurchase godoc
// @Summary    Get a purchase by ID
// @Tags       Purchases
// @Produce    json
// @Param      id             path  int  true  "Purchase ID"
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
		&purchase.Id,
		&purchase.Date,
		&purchase.PaymentDate,
		&purchase.CustomerId,
		&purchase.CreatedAt,
		&purchase.UpdatedAt,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Purchase not found",
		})
		return
	}

	if err := helpers.LinkProductsToPurchase(&purchase, nil); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, purchase)
}

// UpdatePurchase godoc
// @Summary       Update a purchase by ID
// @Tags          Purchases
// @Accept        json
// @Produce       json
// @Param         id       path      int                   true "Purchase ID"
// @Param         purchase body      dtos.EditPurchaseDTO  true "Updated purchase data"
// @Success       200      {object}  models.Purchase
// @Failure       400      {object}  models.ErrorResponse
// @Failure       404      {object}  models.ErrorResponse
// @Failure       500      {object}  models.ErrorResponse
// @Router        /purchases/{id} [put]
func UpdatePurchase(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var input dtos.EditPurchaseDTO
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	var purchase models.Purchase

	err = database.Pool.QueryRow(context.Background(), `
		UPDATE purchases
		SET date = $2
		WHERE id = $1
		RETURNING id, date, payment_date, customer_id, created_at, updated_at
	`,
		id,
		helpers.ParseDate(input.Date),
	).Scan(
		&purchase.Id,
		&purchase.Date,
		&purchase.PaymentDate,
		&purchase.CustomerId,
		&purchase.CreatedAt,
		&purchase.UpdatedAt,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Purchase not found",
		})
		return
	}

	if err := helpers.LinkProductsToPurchase(&purchase, nil); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, purchase)
}

// DeletePurchase godoc
// @Summary       Delete a purchase by ID
// @Tags          Purchases
// @Produce       json
// @Param         id   path      int  true  "Purchase ID"
// @Success       204
// @Failure       400  {object}  models.ErrorResponse
// @Failure       404  {object}  models.ErrorResponse
// @Failure       500  {object}  models.ErrorResponse
// @Router        /purchases/{id} [delete]
func DeletePurchase(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	commandTag, err := database.Pool.Exec(context.Background(), `
		DELETE FROM purchases
		WHERE id = $1
	`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete purchase",
		})
		return
	}

	if commandTag.RowsAffected() == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Purchase not found",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
