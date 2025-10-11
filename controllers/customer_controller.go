package controllers

import (
	"mercadinho/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Customers = []models.Customer{
	{
		ID:           1,
		Name:         "Teste 1",
		Email:        "teste1@example.com",
		Phone:        "5500123456789",
		Address:      "Rua do teste",
		Complement:   nil,
		Number:       "1",
		Neighborhood: "Bairro do teste",
		City:         "Cidade do teste",
		State:        "São Paulo",
		ZIPCode:      "12345678",
		PaymentDay:   15,
	},
}

// GetCustomers godoc
// @Summary      Get all customers
// @Description  Returns a list of all customers
// @Tags         Customers
// @Produce      json
// @Success      200  {array}   models.Customer
// @Failure      404  {object}  models.ErrorResponse
// @Router       /customers [get]
func GetCustomers(context *gin.Context) {
	if len(Customers) == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Nenhum cliente encontrado",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, Customers)
}

// GetCustomer godoc
// @Summary      Get a customer by ID
// @Description  Returns a single customer by ID
// @Tags         Customers
// @Produce      json
// @Param        id   path      int  true  "Customer ID"
// @Success      200  {object}  models.Customer
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /customers/{id} [get]
func GetCustomer(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	for _, customer := range Customers {
		if customer.ID == int(id) {
			context.IndentedJSON(http.StatusOK, customer)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Cliente não encontrado",
	})
}

// CreateCustomer godoc
// @Summary        Create a new customer
// @Description    Adds a new customer to the list
// @Tags           Customers
// @Accept         json
// @Produce        json
// @Param          customer  body      models.Customer  true  "Customer data"
// @Success        201       {object}  models.Customer
// @Failure        400       {object}  models.ErrorResponse
// @Router         /customers [post]
func CreateCustomer(context *gin.Context) {
	var newCustomer models.Customer

	if err := context.BindJSON(&newCustomer); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	Customers = append(Customers, newCustomer)
	context.IndentedJSON(http.StatusCreated, newCustomer)
}
