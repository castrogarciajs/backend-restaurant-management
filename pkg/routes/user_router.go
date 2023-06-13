package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastian009w/backend-restaurant-management/app/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/singup", controllers.SingUp())
	incomingRoutes.POST("/users/login", controllers.Login())
}
