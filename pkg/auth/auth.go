package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Olzheke2003/NewsFeed/internal/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicateUsername = errors.New("username already exists")
)

// RegisterRequest - структура для регистрации пользователя
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"securepassword"`
}

// LoginRequest - структура для входа пользователя
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

type AuthService struct {
	UserRepo *UserRepository
	Secret   string // Секретный ключ для подписи токенов
}

// Создание нового AuthService
func NewAuthService(repo *UserRepository, secret string) *AuthService {
	return &AuthService{
		UserRepo: repo,
		Secret:   secret,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account with provided email, username, and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "User registration data"
// @Success      200 {object} map[string]interface{} "User registered successfully"
// @Failure      400 {object} map[string]string "Invalid request body or missing fields"
// @Failure      409 {object} map[string]string "Email or username already exists"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /auth/register [post]
func (a *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	// Декодируем запрос
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error decoding request:", err) // Логирование ошибки
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверка обязательных полей
	if req.Email == "" || req.Username == "" || req.Password == "" {
		http.Error(w, "Email, username, and password are required", http.StatusBadRequest)
		return
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Создаем пользователя
	user := &models.Users{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	// Сохраняем пользователя в базу данных
	if err := a.UserRepo.CreateUser(user); err != nil {
		if errors.Is(err, ErrDuplicateEmail) || errors.Is(err, ErrDuplicateUsername) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticates a user with email and password and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "User login data"
// @Success      200 {object} map[string]string "JWT token"
// @Failure      400 {object} map[string]string "Invalid request body or missing fields"
// @Failure      401 {object} map[string]string "Invalid email or password"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /auth/login [post]
func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Декодируем запрос
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверка обязательных полей
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Получаем пользователя из базы данных по email
	user, err := a.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		fmt.Println("Error fetching user:", err) // Логирование ошибкиPrintln("Error fetching user:", err) // Логирование ошибки
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Сравниваем хешированный пароль с переданным
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println("Password mismatch:", err) // Логирование ошибки
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Генерируем JWT токен
	token, err := a.generateJWT(user)
	if err != nil {
		fmt.Println("Error generating JWT:", err) // Логирование ошибки
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Возвращаем токен в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Генерация JWT токена
func (a *AuthService) generateJWT(user *models.Users) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Срок действия 24 часа
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.Secret))
}
