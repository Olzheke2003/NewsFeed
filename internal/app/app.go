package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Olzheke2003/NewsFeed/docs"
	"github.com/Olzheke2003/NewsFeed/pkg/auth"

	"github.com/Olzheke2003/NewsFeed/internal/config"
	database "github.com/Olzheke2003/NewsFeed/internal/database/NewsRepository"
	"github.com/Olzheke2003/NewsFeed/internal/services"
	rest "github.com/Olzheke2003/NewsFeed/internal/transport/rest"
	httpSwagger "github.com/swaggo/http-swagger"

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
	s.setupRoutes(db)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) setupRoutes(db *sql.DB) {
	// Создаем экземпляры репозиториев
	newsRepo := database.NewNewsRepository(db)
	userRepo := auth.NewUserRepository(db)

	// Создаем экземпляры сервисов
	newsService := services.NewNewsService(newsRepo)
	authService := auth.NewAuthService(userRepo, s.config.JWTSecret) // Передаем секретный ключ

	// Создаем экземпляры хэндлеров
	newsHandler := rest.NewNewsHandler(newsService)

	// Регистрируем маршруты
	s.router.HandleFunc("/news/comments", newsHandler.GetNewsWithCommentsHandler).Methods("GET")
	s.router.HandleFunc("/news/{id}", newsHandler.GetNewsHandler).Methods("GET")
	s.router.HandleFunc("/auth/register", authService.Register).Methods("POST")
	s.router.HandleFunc("/auth/register", authService.Login).Methods("POST")
	s.router.HandleFunc("/swagger/{any:.*}", httpSwagger.WrapHandler) // Путь для Swagger UI
}
