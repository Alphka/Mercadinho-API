package controllers

import (
	"context"
	"fmt"
	"mercadinho/database"
	"mercadinho/dtos"
	"mercadinho/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary    Get all products
// @Tags       Products
// @Produce    json
// @Success    200  {array}   models.Product
// @Failure    500  {object}  models.ErrorResponse
// @Router     /products [get]
func GetProducts(ctx *gin.Context) {
	rows, err := database.Pool.Query(context.Background(), `
		SELECT id, price, description, created_at, updated_at
		FROM products
	`)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.Id,
			&product.Price,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		products = append(products, product)
	}

	ctx.IndentedJSON(http.StatusOK, products)
}

// GetProduct godoc
// @Summary   Get a product by ID
// @Tags      Products
// @Produce   json
// @Param     id   path      int  true  "Product ID"
// @Success   200  {object}  models.Product
// @Failure   400  {object}  models.ErrorResponse
// @Failure   404  {object}  models.ErrorResponse
// @Router    /products/{id} [get]
func GetProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var product models.Product

	err = database.Pool.QueryRow(context.Background(), `
		SELECT id, price, description, created_at, updated_at
		FROM products
		WHERE id = $1
	`, id).Scan(
		&product.Id,
		&product.Price,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Product not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product   body      dtos.CreateProductDTO  true  "Product data"
// @Success      201       {object}  models.Product
// @Failure      400       {object}  models.ErrorResponse
// @Failure      500       {object}  models.ErrorResponse
// @Router       /products [post]
func CreateProduct(ctx *gin.Context) {
	var newProduct dtos.CreateProductDTO

	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	var product models.Product

	if err := database.Pool.QueryRow(context.Background(), `
		INSERT INTO products (price, description)
		VALUES ($1, $2)
		RETURNING id, price, description, created_at, updated_at
	`,
		newProduct.Price,
		newProduct.Description,
	).Scan(
		&product.Id,
		&product.Price,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create product",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary      Update a product by ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id       path      int                   true "Product ID"
// @Param        product  body      dtos.EditProductDTO   true "Updated product data"
// @Success      200      {object}  models.Product
// @Failure      400      {object}  models.ErrorResponse
// @Failure      404      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /products/{id} [put]
func UpdateProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid ID",
		})
		return
	}

	var input dtos.EditProductDTO
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid data",
		})
		return
	}

	var product models.Product

	err = database.Pool.QueryRow(context.Background(), `
		UPDATE purchases
		SET price = $2,
			description = $3
		WHERE id = $1
		RETURNING id, price, description, created_at, updated_at
	`,
		id,
		input.Price,
		input.Description,
	).Scan(
		&product.Id,
		&product.Price,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Product not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, product)
}
