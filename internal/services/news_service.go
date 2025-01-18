package services

import (
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
func (s *NewsService) GetNewsWithComments() ([]models.NewsWithComments, error) {
	news, err := s.repo.GetAllNews()
	if err != nil {
		return nil, err
	}

	var result []models.NewsWithComments
	for _, n := range news {
		count, err := s.repo.CountComments(n.Title)
		if err != nil {
			return nil, err
		}
		result = append(result, models.NewsWithComments{
			News:          n,
			CommentsCount: count,
		})
	}
	return result, nil
}

func (s *NewsService) GetNews_ID(news_id int) (models.News_id, error) {
	news, err := s.repo.GetNews(news_id)
	if err != nil {
		return news, err
	}
	return news, nil
}
