package main

import (
	"log"

	"github.com/Olzheke2003/NewsFeed/internal/app"
	"github.com/Olzheke2003/NewsFeed/internal/config"
)

func main() {
	cfg := config.New("configs/config.yaml")
	s := app.NewServer(cfg)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
