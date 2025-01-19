package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Olzheke2003/NewsFeed/internal/config" // Путь к вашему пакету конфигурации
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	cfg *config.ServerConfig
	db  *sql.DB
}

func NewAuthHandler(cfg *config.ServerConfig, db *sql.DB) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
		db:  db,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler
// @Summary      Register a new user
// @Description  Creates a new user account with the provided email, username, and password.
// @Tags         auth
// @Accept       json
// @Produce      plain
// @Param        request body RegisterRequest true "User registration data"
// @Success      201 {string} string "User created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error hashing password or creating user"
// @Router       /auth/register [post]
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3)`
	_, err = h.db.Exec(query, req.Email, req.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// LoginHandler
// @Summary      Login a user
// @Description  Authenticates the user and returns a JWT token upon successful login.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "User login data"
// @Success      200 {object} map[string]string "token"
// @Failure      400 {string} string "Invalid input"
// @Failure      401 {string} string "Invalid email or password"
// @Failure      500 {string} string "Database error or error creating token"
// @Router       /auth/login [post]
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var hashedPassword string
	var userID int
	query := `SELECT id, password FROM users WHERE email = $1`
	err := h.db.QueryRow(query, req.Email).Scan(&userID, &hashedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := CheckPassword(hashedPassword, req.Password); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(h.cfg.TokenExpiry).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.cfg.JwtSecretKey))
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
