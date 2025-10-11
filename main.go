package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type customer struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Address      string  `json:"address"`
	Complement   *string `json:"complement"`
	Number       string  `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	ZIPCode      string  `json:"zipCode"`
	PaymentDay   int     `json:"paymentDay"`
}

type purchase struct {
	ID          int        `json:"id"`
	Date        time.Time  `json:"date"`
	TotalPrice  float64    `json:"totalPrice"`
	PaymentDate *time.Time `json:"paymentDate"`
	customerID  int
}

type product struct {
	ID          int     `json:"id"`
	Price       float64 `json:"preco"`
	Description string  `json:"description"`
}

type productPurchase struct {
	ID         int     `json:"id"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	purchaseID int
	productID  int
}

var customers = []customer{
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

var purchases = []purchase{
	{
		ID:          1,
		Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		TotalPrice:  100.00,
		PaymentDate: nil,
		customerID:  1,
	},
}

func main() {
	router := gin.Default()

	// Customers
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomer)
	router.POST("/customers", createCustomer)

	router.Run("localhost:8080")
}

func createCustomer(context *gin.Context) {
	var newCustomer customer

	if err := context.BindJSON(&newCustomer); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	customers = append(customers, newCustomer)
	context.IndentedJSON(http.StatusCreated, newCustomer)
}

func getCustomers(context *gin.Context) {
	if len(customers) == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Nenhum cliente encontrado",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, customers)
}

func getCustomer(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	for _, customer := range customers {
		if customer.ID == int(id) {
			context.IndentedJSON(http.StatusOK, customer)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Cliente não encontrado",
	})
}
