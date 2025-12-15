package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/henlan93/go-crud-todo/db"
	"github.com/henlan93/go-crud-todo/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to the database
	if err := db.Init(); err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Setup router
	r := routes.SetupRouter()

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Graceful shutdown setup
	serverErr := make(chan error)
	go func() {
		log.Printf("Server running on port %s", port)
		serverErr <- r.Run(":" + port)
	}()

	// Listen for interrupt signals (CTRL+C, Docker stop, etc.)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		log.Fatalf("Server error: %v", err)

	case <-quit:
		log.Println("Shutting down gracefully...")
	}

	log.Println("Server stopped")
}
