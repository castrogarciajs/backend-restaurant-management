package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func SingUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(password, provider string) (bool, string) {
	return true, ""
}
