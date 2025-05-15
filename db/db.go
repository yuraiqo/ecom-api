package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPostgreSQLStorage(connStr string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), connStr)

	return db, err
}
