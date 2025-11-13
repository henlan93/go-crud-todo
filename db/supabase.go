package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Init() error {
	var err error
	Conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	return err
}
