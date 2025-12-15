package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/henlan93/go-crud-todo/handlers"
	"github.com/henlan93/go-crud-todo/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	todos := r.Group("/todos", middleware.AuthMiddleware())
	{
		todos.POST("/", handlers.CreateTodo)
		todos.GET("/:id", handlers.GetTodo)
		todos.PUT("/:id", handlers.UpdateTodo)
		todos.DELETE("/:id", handlers.DeleteTodo)
	}

	return r
}
