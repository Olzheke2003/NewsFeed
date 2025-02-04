package services

import (
	"log"

	database "github.com/Olzheke2003/NewsFeed/internal/database/NewsRepository"
	"github.com/Olzheke2003/NewsFeed/internal/models"
)

type NewsService struct {
	repo *database.NewsRepository
}

func NewNewsService(repo *database.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

// Получить все новости с количеством комментариев
func (s *NewsService) GetNewsWithComments() ([]models.News, error) {
	news, err := s.repo.GetAllNews()
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (s *NewsService) GetNews_ID(news_id int) (models.News_id, error) {
	news, err := s.repo.GetNews(news_id)
	if err != nil {
		return news, err
	}
	return news, nil
}

func (s *NewsService) DeleteNewsService(news_id int) error {
	news := s.repo.DeleteNews(news_id)
	if news != nil {
		log.Printf("Error")
		return news
	}
	return nil
}


func (s *NewsService) UpdateNewsService(newsID int, title string, content string, image string) error {
	news := s.repo.UpdateNews(newsID, title, content, image)
	if news != nil {
		log.Printf("Error")
		return news
	}
	return nil
}