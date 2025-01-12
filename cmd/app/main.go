package main

import (
	"log"

	"github.com/Olzheke2003/NewsFeed/internal/app"
	"github.com/Olzheke2003/NewsFeed/internal/config"
)

// @title NewsFeed API
// @version 1.0
// @description API для управления новостями и комментариями.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.New("configs/config.yaml")
	s := app.NewServer(cfg)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
