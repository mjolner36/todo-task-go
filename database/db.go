package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var DB *pgxpool.Pool

func InitDatabase() {
	var err error
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database %v", err)
	}
	fmt.Println("Connected to database %v", DB)
}
