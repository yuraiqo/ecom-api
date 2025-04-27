package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/yuraiqo/ecom/service/user"
)

type APIServer struct {
	port string
	db   *sql.DB
}

func NewAPIServer(port string, db *sql.DB) *APIServer {
	return &APIServer{port: port, db: db}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.port)

	return http.ListenAndServe(s.port, router)
}
