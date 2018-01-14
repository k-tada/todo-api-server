package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"todo-api-server/routes/todos"
	// "github.com/k-tada/todo-api-server/routes/todos"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "access to /todos"})
	})
	r.GET("/todos", todos.List)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9999"
	}
	r.Run(":" + port)
}
