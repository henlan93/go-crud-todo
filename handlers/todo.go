package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/henlan93/go-crud-todo/db"
	"github.com/henlan93/go-crud-todo/models"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO todos (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, title, completed FROM todos WHERE id=$1", id).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, _ := strconv.Atoi(id)
	todo.ID = uid

	_, err := db.Conn.Exec(context.Background(),
		"UPDATE todos SET title=$1, completed=$2 WHERE id=$3", todo.Title, todo.Completed, todo.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM todos WHERE id=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
