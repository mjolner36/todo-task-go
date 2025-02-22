package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func InitDatabase(connString string) {
	var err error
	DB, err = pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database %v", err)
	}
	fmt.Println("Connected to database %v", DB)
}
