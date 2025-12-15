package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

// Init establishes a database connection using the DATABASE_URL environment variable.
func Init() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	Conn = conn
	log.Println("✅ Database connection established")
	return nil
}

// Close gracefully closes the database connection.
func Close() {
	if Conn != nil {
		_ = Conn.Close(context.Background())
		log.Println("✅ Database connection closed")
	}
}
