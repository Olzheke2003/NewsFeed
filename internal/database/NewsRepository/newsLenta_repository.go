package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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
		SELECT n.title, n.image, COUNT(c.id) AS comments_count
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
		if err := rows.Scan(&n.Title, &n.Image, &n.CommentsCount); err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	return news, nil
}

func (r *NewsRepository) GetNews(newsID int) (models.News_id, error) {
	var news models.News_id

	// Запрос для получения новости с комментариями в формате JSON
	rows, err := r.DB.Query(`
        SELECT
            n.id,
            n.title,
            n.content,
            n.image,
            json_agg(
                json_build_object(
                    'content', c.content
                )
            ) AS comments
        FROM news n
        LEFT JOIN comments c ON n.id = c.news_id
        WHERE n.id = $1
        GROUP BY n.id
    `, newsID)
	if err != nil {
		log.Printf("Failed to execute query for newsID %d: %v", newsID, err)
		return models.News_id{}, err
	}
	defer rows.Close()

	// Чтение результата
	if rows.Next() {
		var commentsJSON []byte // Тип для хранения JSON данных
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Image, &commentsJSON)
		if err != nil {
			log.Printf("Failed to scan news data for newsID %d: %v", newsID, err)
			return models.News_id{}, err
		}

		// Распарсим JSON в структуру Comment
		err = json.Unmarshal(commentsJSON, &news.Comments)
		if err != nil {
			log.Printf("Failed to unmarshal comments JSON for newsID %d: %v", newsID, err)
			return models.News_id{}, err
		}
	} else {
		log.Printf("No news found for newsID %d", newsID)
		return models.News_id{}, fmt.Errorf("news with id %d not found", newsID)
	}

	// Проверка на ошибки после завершения итерации
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows for newsID %d: %v", newsID, err)
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

func (r *NewsRepository) UpdateNews(newsID int, title, content, image string) error {
	query := "UPDATE news SET title = $1, content = $2, image = $3 WHERE id = $4"
	result, err := r.DB.Exec(query, title, content, image, newsID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении новости: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении количества изменённых строк: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("новость с id %d не найдена", newsID)
	}

	return nil
}
