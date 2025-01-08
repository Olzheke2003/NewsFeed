package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Olzheke2003/NewsFeed/internal/config"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Server struct {
	config *config.ServerConfig
	logger *log.Logger
	router *mux.Router
}

func NewServer(cfg *config.ServerConfig) *Server {
	router := mux.NewRouter()
	return &Server{
		config: cfg,
		logger: log.Default(),
		router: router,
	}
}

func (s *Server) Run() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	defer db.Close()

	s.logger.Println("Starting server on", s.config.BindAddr)
	s.setupRoutes()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to NewsFeed!")
	}).Methods("GET")
}
