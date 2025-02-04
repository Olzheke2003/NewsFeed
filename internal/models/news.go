package models

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
	Title         string `json:"title"`
	Image         string `json:"image"`
	CommentsCount int    `json:"comments_count"`
}

type News_id struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
	Comments []Comment `json:"comments"` // Массив комментариев
}

type Comment struct {
	Content string `json:"content"`
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

type DeleteNews struct {
	ID int `json:"id"`
}

// ErrorResponse описывает структуру для ошибок API
type ErrorResponse struct {
	Error string `json:"error"`
}

