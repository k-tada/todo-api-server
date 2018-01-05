package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World123",
		})
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9999"
	}
	r.Run(":" + port)
}
