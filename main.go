package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")

	r := gin.New()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if PORT == "" {
		PORT = "8000"
	}
	r.Run()
}
