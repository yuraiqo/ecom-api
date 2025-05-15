package api

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/yuraiqo/ecom/service/user"
)

type APIServer struct {
	port string
	db   *pgx.Conn
}

func NewAPIServer(port string, db *pgx.Conn) *APIServer {
	return &APIServer{port: port, db: db}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/health", s.healthCheckHandler)

	v1 := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	log.Println("Listening on", s.port)

	return http.ListenAndServe(s.port, router)
}

func (s *APIServer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is healthy"))
}
