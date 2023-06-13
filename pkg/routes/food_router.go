package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/app/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/food/:food_id", controllers.GetFood())
	incomingRoutes.POST("foods", controllers.CreateFood())
	incomingRoutes.PATCH("food/:food_id", controllers.UpdateFood())

}
