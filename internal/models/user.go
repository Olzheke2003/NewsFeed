package models

import (
	"database/sql"
	"encoding/json"
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

func (u *Users) MarshalJSON() ([]byte, error) {
	type Alias Users
	return json.Marshal(&struct {
		Image       interface{} `json:"image"`
		Name        interface{} `json:"name"`
		Lastname    interface{} `json:"lastname"`
		DateOfBirth interface{} `json:"date_of_birth"`
		GenderID    interface{} `json:"gender_id"`
		*Alias
	}{
		Image:    u.Image.String,
		Name:     u.Name.String,
		Lastname: u.Lastname.String,
		DateOfBirth: func() interface{} {
			if u.DateOfBirth.Valid {
				return u.DateOfBirth.Time.Format("2006-01-02") // Преобразуем в строку в формате YYYY-MM-DD
			}
			return nil // Если NULL, то возвращаем nil
		}(),
		GenderID: func() interface{} {
			if u.GenderID.Valid {
				return u.GenderID.Int64 // Если есть значение, возвращаем его
			}
			return nil // Если NULL, возвращаем nil
		}(),
		Alias: (*Alias)(u),
	})
}
