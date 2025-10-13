package controllers

import (
	"context"
	"mercadinho/database"
	"mercadinho/dtos"
	"mercadinho/helpers"
	"mercadinho/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetCustomers godoc
// @Summary     Get all customers
// @Tags        Customers
// @Produce     json
// @Success     200  {array}   models.Customer
// @Failure     500  {object}  models.ErrorResponse
// @Router      /customers [get]
func GetCustomers(ctx *gin.Context) {
	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day, created_at, updated_at
		FROM customers
	`)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	defer rows.Close()

	customers := []models.Customer{}

	for rows.Next() {
		var customer models.Customer

		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.Address,
			&customer.Number,
			&customer.Complement,
			&customer.Neighborhood,
			&customer.City,
			&customer.State,
			&customer.ZIPCode,
			&customer.PaymentDay,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		customers = append(customers, customer)
	}

	ctx.IndentedJSON(http.StatusOK, customers)
}

// GetCustomer godoc
// @Summary    Get a customer by ID
// @Tags       Customers
// @Produce    json
// @Param      id   path      int  true  "Customer ID"
// @Success    200  {object}  models.Customer
// @Failure    400  {object}  models.ErrorResponse
// @Failure    404  {object}  models.ErrorResponse
// @Router     /customers/{id} [get]
func GetCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var customer models.Customer

	err = database.Pool.QueryRow(context.Background(), `
		SELECT id, name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day, created_at, updated_at
		FROM customers
		WHERE id = $1
	`, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
		&customer.Number,
		&customer.Complement,
		&customer.Neighborhood,
		&customer.City,
		&customer.State,
		&customer.ZIPCode,
		&customer.PaymentDay,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Customer not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, customer)
}

// CreateCustomer godoc
// @Summary       Create a new customer
// @Tags          Customers
// @Accept        json
// @Produce       json
// @Param         customer  body      dtos.CreateCustomerDTO  true  "Customer data"
// @Success       201       {object}  models.Customer
// @Failure       400       {object}  models.ErrorResponse
// @Failure       500       {object}  models.ErrorResponse
// @Router        /customers [post]
func CreateCustomer(ctx *gin.Context) {
	var newCustomer dtos.CreateCustomerDTO

	if err := ctx.BindJSON(&newCustomer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	var customer models.Customer

	if err := database.Pool.QueryRow(context.Background(), `
		INSERT INTO customers (name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day, created_at, updated_at
	`,
		newCustomer.Name,
		newCustomer.Email,
		newCustomer.Phone,
		newCustomer.Address,
		newCustomer.Number,
		newCustomer.Complement,
		newCustomer.Neighborhood,
		newCustomer.City,
		newCustomer.State,
		newCustomer.ZIPCode,
		newCustomer.PaymentDay,
	).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
		&customer.Number,
		&customer.Complement,
		&customer.Neighborhood,
		&customer.City,
		&customer.State,
		&customer.ZIPCode,
		&customer.PaymentDay,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create customer",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, customer)
}

// UpdateCustomer godoc
// @Summary       Update a customer by ID
// @Tags          Customers
// @Accept        json
// @Produce       json
// @Param         id       path      int                   true "Customer ID"
// @Param         customer body      dtos.EditCustomerDTO  true "Updated customer data"
// @Success       200      {object}  models.Customer
// @Failure       400      {object}  models.ErrorResponse
// @Failure       404      {object}  models.ErrorResponse
// @Failure       500      {object}  models.ErrorResponse
// @Router        /customers/{id} [put]
func UpdateCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var input dtos.EditCustomerDTO
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	var customer models.Customer

	err = database.Pool.QueryRow(context.Background(), `
		UPDATE customers
		SET name = $2,
			email = $3,
			phone = $4,
			address = $5,
			complement = $6,
			number = $7,
			neighborhood = $8,
			city = $9,
			state = $10,
			zip_code = $11,
			payment_day = $12
		WHERE id = $1
		RETURNING id, name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day, created_at, updated_at
	`,
		id,
		input.Name,
		input.Email,
		input.Phone,
		input.Address,
		input.Number,
		input.Complement,
		input.Neighborhood,
		input.City,
		input.State,
		input.ZIPCode,
		input.PaymentDay,
	).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
		&customer.Number,
		&customer.Complement,
		&customer.Neighborhood,
		&customer.City,
		&customer.State,
		&customer.ZIPCode,
		&customer.PaymentDay,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Customer not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, customer)
}

// DeleteCustomer godoc
// @Summary       Delete a customer by ID
// @Tags          Customers
// @Produce       json
// @Param         id   path      int  true  "Customer ID"
// @Success       204
// @Failure       400  {object}  models.ErrorResponse
// @Failure       404  {object}  models.ErrorResponse
// @Failure       500  {object}  models.ErrorResponse
// @Router        /customers/{id} [delete]
func DeleteCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	commandTag, err := database.Pool.Exec(context.Background(), `
		DELETE FROM customers
		WHERE id = $1
	`, id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to delete customer",
		})
		return
	}

	if commandTag.RowsAffected() == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Customer not found",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetCustomerPurchases godoc
// @Summary             Get all customer purchases
// @Tags                Customers
// @Produce             json
// @Param               id   path      int  true  "Customer ID"
// @Success             200  {array}   models.Purchase
// @Failure             500  {object}  models.ErrorResponse
// @Router              /customers/{id}/purchases [get]
func GetCustomerPurchases(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, date, payment_date, customer_id, created_at, updated_at
		FROM purchases
		WHERE customer_id = $1
	`, customerId)
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

	ctx.IndentedJSON(http.StatusOK, purchases)
}

// CreateCustomerPurchase godoc
// @Summary               Register a new purchase
// @Tags                  Customers
// @Accept                json
// @Produce               json
// @Param                 id        path      int                     true  "Customer ID"
// @Param                 purchase  body      dtos.CreatePurchaseDTO  true  "Purchase data"
// @Success               201       {object}  models.Purchase
// @Failure               400       {object}  models.ErrorResponse
// @Failure               500       {object}  models.ErrorResponse
// @Router                /customers/{id}/purchases [post]
func CreateCustomerPurchase(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var newPurchase dtos.CreatePurchaseDTO

	if err := ctx.BindJSON(&newPurchase); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	tx, err := database.Pool.Begin(context.Background())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to start transaction",
		})
		return
	}

	defer tx.Rollback(context.Background())

	var totalPrice float64

	if err := tx.QueryRow(context.Background(), "SELECT COALESCE(SUM(price), 0) FROM products WHERE id = ANY($1)", newPurchase.Products).Scan(&totalPrice); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to calculate total price",
		})
		return
	}

	var paymentDate *time.Time
	if newPurchase.PaymentDate != nil {
		paymentDate = helpers.ParseDate(*newPurchase.PaymentDate)
	} else {
		paymentDate = nil
	}

	var purchase models.Purchase

	err = tx.QueryRow(context.Background(), `
		INSERT INTO purchases (customer_id, total_price, date, payment_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, customer_id, total_price, date, payment_date, created_at, updated_at
	`,
		customerId,
		totalPrice,
		helpers.ParseDate(newPurchase.Date),
		paymentDate,
	).Scan(
		&purchase.Id,
		&purchase.CustomerId,
		&purchase.TotalPrice,
		&purchase.Date,
		&purchase.PaymentDate,
		&purchase.CreatedAt,
		&purchase.UpdatedAt,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to register purchase",
		})
		return
	}

	for _, productId := range newPurchase.Products {
		_, err = tx.Exec(context.Background(), `
			INSERT INTO product_purchase (product_id, purchase_id)
			VALUES ($1, $2)
		`, productId, purchase.Id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Failed to link products to purchase",
			})
			return
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to commit transaction",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, purchase)
}
