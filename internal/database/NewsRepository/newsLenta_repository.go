package database

import (
	"database/sql"
	"fmt"

	"github.com/Olzheke2003/NewsFeed/internal/models"
)

type NewsRepository struct {
	DB *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{DB: db}
}

// Получить все новости
func (r *NewsRepository) GetAllNews() ([]models.News, error) {
	rows, err := r.DB.Query(`
		SELECT n.title, n.created_at, n.image, COUNT(c.id) AS comments_count
		FROM news n
		LEFT JOIN 
			comments c ON n.id = c.news_id
		GROUP BY 
			n.id
		ORDER BY 
			n.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.News
	for rows.Next() {
		var n models.News
		if err := rows.Scan(&n.Title, &n.CreatedAt, &n.Image, &n.CommentsCount); err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}

func (r *NewsRepository) GetNews(newsID int) (models.News_id, error) {
	var news models.News_id
	err := r.DB.QueryRow("SELECT id, title, created_at, content, image FROM news WHERE id = $1", newsID).
		Scan(&news.ID, &news.Title, &news.CreatedAt, &news.Content, &news.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.News_id{}, fmt.Errorf("news with id %d not found", newsID)
		}
		return models.News_id{}, err
	}
	return news, nil
}

// Удалить новость по ID
func (r *NewsRepository) DeleteNews(newsID int) error {
	result, err := r.DB.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no news found with id %d", newsID)
	}

	return nil
}
