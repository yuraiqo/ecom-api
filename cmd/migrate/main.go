package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/yuraiqo/ecom/config"
)

func main() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.Host,
		config.Envs.Port,
		config.Envs.DBName,
	)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"ecom",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
