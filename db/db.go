package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgreSQLStorage(connStr string) *pgx.Conn {
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
