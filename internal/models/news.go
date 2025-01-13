package models

import "time"

// type Like struct {
// 	ID         int       `json:"id"`
// 	UserID     int       `json:"user_id"`     // ID пользователя, который поставил лайк
// 	ObjectID   int       `json:"object_id"`   // ID объекта (новости или комментария)
// 	ObjectType string    `json:"object_type"` // Тип объекта: "news" или "comment"
// 	CreatedAt  time.Time `json:"created_at"`  // Время, когда был поставлен лайк
// }

type Category struct {
	ID       int    `json:"id"`
	Category string `json:"category_id"`
}

type News struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Image     string    `json:"image"`
}

type News_id struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
}

type Comments struct {
	ID       int        `json:"id"`
	NewsID   int        `json:"news_id"`
	ParentID *int       `json:"parent_id"`
	UserID   int        `json:"user_id"`
	Content  string     `json:"content"`
	DatePost string     `json:"date_post"`
	Replies  []Comments `json:"replies"`
}
