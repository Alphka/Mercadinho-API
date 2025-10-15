package controllers

import (
	"context"
	"fmt"
	"mercadinho/database"
	"mercadinho/dtos"
	"mercadinho/helpers"
	"mercadinho/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetCustomers godoc
// @Summary     Get all customers
// @Tags        Customers
// @Produce     json
// @Param       page           query  int  false  "Page number" default(1)
// @Param       pageSize       query  int  false  "Number of items per page" default(20)
// @Success     200  {array}   models.PaginatedResponse[models.Customer]
// @Failure     500  {object}  models.ErrorResponse
// @Router      /customers [get]
func GetCustomers(ctx *gin.Context) {
	var totalItems int

	err := database.Pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM products`).Scan(&totalItems)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	page, pageSize, totalPages := helpers.ValidatePagination(ctx, totalItems)

	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, name, email, phone, address, number, complement, neighborhood, city, state, zip_code, payment_day, created_at, updated_at
		FROM customers
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

	ctx.IndentedJSON(http.StatusOK, models.PaginatedResponse[models.Customer]{
		Data:       customers,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	})
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

// GetCustomerPendingPayments godoc
// @Summary               Get all customer purchases with pending payments
// @Tags                  Customers
// @Produce               json
// @Param                 id        path      int                     true  "Customer ID"
// @Success               200       {object}  []models.CustomerPendingPayment
// @Failure               400       {object}  models.ErrorResponse
// @Failure               500       {object}  models.ErrorResponse
// @Router                /customers/{id}/pending-payments [get]
func GetCustomerPendingPayments(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, date, total_price, created_at, updated_at
		FROM purchases
		WHERE customer_id = $1 AND payment_date IS NULL
		ORDER BY created_at DESC
	`, customerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	defer rows.Close()

	purchases := []models.CustomerPendingPayment{}

	for rows.Next() {
		var purchase models.CustomerPendingPayment

		err := rows.Scan(
			&purchase.Id,
			&purchase.Date,
			&purchase.TotalPrice,
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

// GetCustomerOverduePayments godoc
// @Summary               Get all customer purchases that are overdue (past payment date)
// @Tags                  Customers
// @Produce               json
// @Param                 id        path      int  true  "Customer ID"
// @Success               200       {object}  []models.CustomerPendingPayment
// @Failure               400       {object}  models.ErrorResponse
// @Failure               500       {object}  models.ErrorResponse
// @Router                /customers/{id}/overdue-payments [get]
func GetCustomerOverduePayments(ctx *gin.Context) {
	customerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	rows, err := database.Pool.Query(context.Background(), `
		SELECT product.id, product.date, product.total_price, product.created_at, product.updated_at
		FROM purchases product
		JOIN customers customer ON customer.id = product.customer_id
		WHERE
			product.customer_id = $1
			AND product.payment_date IS NULL
			AND CURRENT_DATE > (
				date_trunc('month', product.date) + interval '1 month' + (customer.payment_day - 1) * interval '1 day'
			)
		ORDER BY product.created_at DESC
	`, customerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	defer rows.Close()

	purchases := []models.CustomerPendingPayment{}

	for rows.Next() {
		var purchase models.CustomerPendingPayment

		err := rows.Scan(
			&purchase.Id,
			&purchase.Date,
			&purchase.TotalPrice,
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

// GetCustomerPurchases godoc
// @Summary             Get all customer purchases
// @Tags                Customers
// @Produce             json
// @Param               id             path  int  true  "Customer ID"
// @Param               page           query  int  false  "Page number" default(1)
// @Param               pageSize       query  int  false  "Number of items per page" default(20)
// @Success             200  {array}   models.PaginatedResponse[models.Purchase]
// @Failure             400  {object}  models.ErrorResponse
// @Failure             404  {object}  models.ErrorResponse
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

	if err := database.Pool.QueryRow(context.Background(), "SELECT id FROM customers WHERE id = $1", customerId).Scan(&customerId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Customer not found",
		})
		return
	}

	var totalItems int

	err = database.Pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM purchases`).Scan(&totalItems)
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
		WHERE customer_id = $1
		ORDER BY created_at DESC
		OFFSET (($2 - 1) * $3)
		LIMIT $3
	`, customerId, page, pageSize)
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

// CreateCustomerPurchase godoc
// @Summary               Register a new purchase
// @Tags                  Customers
// @Accept                json
// @Produce               json
// @Param                 id        path      int                     true  "Customer ID"
// @Param                 purchase  body      dtos.CreatePurchaseDTO  true  "Purchase data"
// @Success               201       {object}  models.Purchase
// @Failure               400       {object}  models.ErrorResponse
// @Failure               404       {object}  models.ErrorResponse
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

	if err := database.Pool.QueryRow(context.Background(), "SELECT id FROM customers WHERE id = $1", customerId).Scan(&customerId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Customer not found",
		})
		return
	}

	productIds := make([]int, len(newPurchase.Products))
	for index, product := range newPurchase.Products {
		productIds[index] = product.Id
	}

	rows, err := database.Pool.Query(context.Background(), "SELECT id, price FROM products WHERE id = ANY($1)", productIds)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to check products",
		})
		return
	}
	defer rows.Close()

	foundProductsWithPrice := make(map[int]float64)

	for rows.Next() {
		var id int
		var price float64

		if err := rows.Scan(&id, &price); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Failed reading product rows",
			})
			return
		}

		foundProductsWithPrice[id] = price
	}

	var productsNotFound []string
	var totalPrice float64

	for _, product := range newPurchase.Products {
		price, exists := foundProductsWithPrice[product.Id]

		if !exists {
			productsNotFound = append(productsNotFound, strconv.Itoa(product.Id))
			continue
		}

		totalPrice += price * float64(product.Quantity)
	}

	if len(productsNotFound) > 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: fmt.Sprintf("Products not found: %s", strings.Join(productsNotFound, ", ")),
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

	for _, product := range newPurchase.Products {
		price := foundProductsWithPrice[product.Id]

		_, err = tx.Exec(context.Background(), `
			INSERT INTO product_purchase (price, quantity, product_id, purchase_id)
			VALUES ($1, $2, $3, $4)
		`, price, product.Quantity, product.Id, purchase.Id)
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

	var foundProductIds []int

	for key := range foundProductsWithPrice {
		foundProductIds = append(foundProductIds, key)
	}

	if err := helpers.LinkProductsToPurchase(&purchase, &foundProductIds); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, purchase)
}
