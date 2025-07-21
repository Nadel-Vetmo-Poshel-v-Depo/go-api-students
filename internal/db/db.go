package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDB() *pgx.Conn {
	connStr := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") + "?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	return conn
}
