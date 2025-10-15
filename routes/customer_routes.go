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
			customers.PUT("/:id", controllers.UpdateCustomer)
			customers.DELETE("/:id", controllers.DeleteCustomer)
			customers.GET("/:id/pending-payments", controllers.GetCustomerPendingPayments)
			customers.GET("/:id/overdue-payments", controllers.GetCustomerOverduePayments)

			customers.GET("/:id/purchases", controllers.GetCustomerPurchases)
			customers.POST("/:id/purchases", controllers.CreateCustomerPurchase)
		}
	}
}
