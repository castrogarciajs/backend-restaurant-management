package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/app/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:food_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:food_id", controllers.UpdateInvoice())

}
