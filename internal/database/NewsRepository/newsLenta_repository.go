package database

import (
	"database/sql"
	"myapp/internal/models"
)

type NewsRepository struct {
	DB *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
	return &NewsRepository{DB: db}
}

// Получить все новости
func (r *NewsRepository) GetAllNews() ([]models.News, error) {
	rows, err := r.DB.Query("SELECT id, title, category_id, created_at, content, image FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.News
	for rows.Next() {
		var n models.News
		if err := rows.Scan(&n.ID, &n.Title, &n.CategoryID, &n.CreatedAt, &n.Content, &n.Image); err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}

// Подсчитать количество комментариев для новости
func (r *NewsRepository) CountComments(newsID int) (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM comments WHERE news_id = $1", newsID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
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
