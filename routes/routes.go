package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henlan93/go-crud-todo/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	todos := r.Group("/todos")
	{
		todos.POST("/", handlers.CreateTodo)
		todos.GET("/:id", handlers.GetTodo)
		todos.PUT("/:id", handlers.UpdateTodo)
		todos.DELETE("/:id", handlers.DeleteTodo)
	}

	return r
}
