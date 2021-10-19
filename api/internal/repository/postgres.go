package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func NewPostgresDB() (*pgx.Conn, error) {

	pgDBName := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgUser, pgPass, "postgresql", "5432", pgDBName, "disable")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You connect to your database")
	return conn, err
}
