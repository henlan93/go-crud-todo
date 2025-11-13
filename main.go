package main

import (
	"log"
	"os"

	"github.com/henrilan93/go-crud-todo/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
