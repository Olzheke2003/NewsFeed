package models

import "time"

type Category struct {
	ID            int    `json:id`
	Category_name string `json:category_name`
}

type News struct {
	ID         int       `json:id`
	Title      string    `json:title`
	CategoryID Category  `json:category_id`
	CreatedAt  time.Time `json:created_at`
	Content    string    `json:"content"`
	Image      string    `json:"image"`
	Likes      int       `json:likes`
}

type Comments struct {
	ID       int       `json:"id"`
	NewsID   News      `json:"news_id"`
	ParentID *int      `json:"parent_id"`
	UserID   Users     `json:"user_id"`
	Content  string    `json:"content"`
	DatePost string    `json:"date_post"`
	Replies  []Comment `json:"replies"`
	Likes    int       `json:likes`
}
