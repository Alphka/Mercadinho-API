package routes

import (
	"mercadinho/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customers.GET("", controllers.GetCustomers)
			customers.GET("/:id", controllers.GetCustomer)
			customers.POST("", controllers.CreateCustomer)
		}
	}
}
