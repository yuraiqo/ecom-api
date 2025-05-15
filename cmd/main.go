package main

import (
	"fmt"
	"log"

	"github.com/yuraiqo/ecom/cmd/api"
	"github.com/yuraiqo/ecom/config"
	"github.com/yuraiqo/ecom/db"
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

	db, err := db.NewPostgreSQLStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
