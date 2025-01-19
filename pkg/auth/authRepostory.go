package auth

import (
	"database/sql"
	"time"

	"github.com/Olzheke2003/NewsFeed/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Сохранить пользователя в базу данных
func (r *UserRepository) CreateUser(user *models.Users) error {
	// Проверяем уникальность email
	var existingID int
	err := r.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingID)
	if err == nil {
		return ErrDuplicateEmail
	}
	if err != sql.ErrNoRows {
		return err
	}

	// Проверяем уникальность username
	err = r.DB.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&existingID)
	if err == nil {
		return ErrDuplicateUsername
	}
	if err != sql.ErrNoRows {
		return err
	}

	// Вставляем пользователя
	query := `
		INSERT INTO users (email, username, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	return r.DB.QueryRow(query, user.Email, user.Username, user.Password, time.Now()).Scan(&user.ID)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.Users, error) {
	query := `SELECT id, email, username, password, image, name, lastname, date_of_birth, gender_id, created_at FROM users WHERE email = $1`
	var user models.Users
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Image,
		&user.Name,
		&user.Lastname,
		&user.DateOfBirth,
		&user.GenderID,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
