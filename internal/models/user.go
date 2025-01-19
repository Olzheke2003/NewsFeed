package models

import (
	"database/sql"
	"time"
)

type Gender struct {
	ID     int    `json:"id"`
	Gender string `json:gender`
}

type Users struct {
	ID          int            `json:"id"`
	Email       string         `json:"email"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Image       sql.NullString `json:"image"`         // Используем sql.NullString для возможного NULL
	Name        sql.NullString `json:"name"`          // Используем sql.NullString
	Lastname    sql.NullString `json:"lastname"`      // Используем sql.NullString
	DateOfBirth sql.NullTime   `json:"date_of_birth"` // Используем sql.NullTime
	GenderID    sql.NullInt64  `json:"gender_id"`     // Используем sql.NullInt64 для числовых полей, которые могут быть NULL
	CreatedAt   time.Time      `json:"created_at"`
}
