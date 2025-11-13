package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henlan93/go-crud-todo/db"
	"github.com/henlan93/go-crud-todo/handlers"
)

func SetupRouter() *gin.Engine {
	_ = db.Init()

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/todos", handlers.CreateTodo)
		api.GET("/todos/:id", handlers.GetTodo)
		api.PUT("/todos/:id", handlers.UpdateTodo)
		api.DELETE("/todos/:id", handlers.DeleteTodo)
	}
	return r
}
