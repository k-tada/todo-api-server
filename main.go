package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"todo-api-server/db"
	"todo-api-server/middlewares"
	"todo-api-server/routes/todos"
	// "github.com/k-tada/todo-api-server/routes/todos"
)

func init() {
	db.Connect()
}

func main() {
	r := gin.Default()

	// Middlewares
	r.Use(middlewares.Connect)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "access to /todos"})
	})

	todosRoute := r.Group("/todos")
	{
		todosRoute.GET("/", todos.List)
		todosRoute.POST("/", todos.Create)
		todosRoute.GET("/:_id", todos.Show)
		todosRoute.PUT("/:_id", todos.Update)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9999"
	}
	r.Run(":" + port)
}
