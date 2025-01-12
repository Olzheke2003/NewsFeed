package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Olzheke2003/NewsFeed/docs"

	"github.com/Olzheke2003/NewsFeed/internal/config"
	"github.com/Olzheke2003/NewsFeed/internal/database/NewsRepository"
	"github.com/Olzheke2003/NewsFeed/internal/services"
	rest "github.com/Olzheke2003/NewsFeed/internal/transport/rest" // Правильный импорт rest

	// Импортируйте ваш сгенерированный Swagger файл
	"github.com/swaggo/http-swagger" // Импортируйте этот пакет

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
	// Создаем экземпляр репозитория, передаем подключение к базе данных
	newsRepo := database.NewNewsRepository(db)

	// Создаем экземпляр сервиса
	newsService := services.NewNewsService(newsRepo)

	// Создаем экземпляр хэндлера с использованием rest
	newsHandler := rest.NewNewsHandler(newsService)

	// Регистрируем маршрут
	s.router.HandleFunc("/news/comments", newsHandler.GetNewsWithCommentsHandler).Methods("GET")
	s.router.HandleFunc("/swagger/{any:.*}", httpSwagger.WrapHandler) // Путь для Swagger UI
}
