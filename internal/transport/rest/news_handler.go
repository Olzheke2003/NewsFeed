package transport

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	
	"github.com/Olzheke2003/NewsFeed/internal/services"
	"github.com/gorilla/mux"
)

type NewsUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type NewsHandler struct {
	service *services.NewsService
}

func NewNewsHandler(service *services.NewsService) *NewsHandler {
	return &NewsHandler{service: service}
}

// GetNewsWithCommentsHandler возвращает список новостей с количеством комментариев.
// @Summary Get news with comments
// @Description Get all news with their respective comment counts
// @Tags news
// @Accept json
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {object} map[string]string
// @Router /news/comments [get]
func (h *NewsHandler) GetNewsWithCommentsHandler(w http.ResponseWriter, r *http.Request) {
	news, err := h.service.GetNewsWithComments()
	if err != nil {
		http.Error(w, "Failed to get news with comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// GetNewsHandler возвращает одну новость по ID.
// @Summary Get news by ID
// @Description Get a single news item by its ID
// @Tags news
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} models.News
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /news/{id} [get]
func (h *NewsHandler) GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL
	vars := mux.Vars(r) // Используется для работы с gorilla/mux
	newsIDParam := vars["id"]

	// Преобразуем ID из строки в int
	newsID, err := strconv.Atoi(newsIDParam)
	if err != nil {
		http.Error(w, `{"error": "Invalid news ID"}`, http.StatusBadRequest)
		return
	}

	// Получаем новость из сервиса
	news, err := h.service.GetNews_ID(newsID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, `{"error": "News not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error": "Failed to retrieve news"}`, http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// DeleteNews godoc
// @Summary Delete a news article
// @Description Delete a news article by ID
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path int true "News ID"
// @Success 200 {string} string "GOOD DELETE"
// @Failure 400 {object} models.ErrorResponse "Invalid news ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /news/{id} [delete]
func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsIDParam := vars["id"]
	newsID, err := strconv.Atoi(newsIDParam)
	if err != nil {
		http.Error(w, `{"error": "Invalid news ID"}`, http.StatusBadRequest)
		return
	}
	if h.service.DeleteNewsService(newsID) != nil {
		http.Error(w, `{"error": "Failed to delete news"}`, http.StatusBadRequest)
		return
	} else {
		http.Error(w, `"GOOD DELETE"`, http.StatusOK)
		return
	}
}



// UpdateNews godoc
// @Summary Update a news article
// @Description Update the title, content, and image of a news article
// @Tags news
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Param news body NewsUpdateRequest true "Updated news data"
// @Success 200 {string} string "News updated successfully"
// @Failure 400 {object} map[string]string "Invalid request data"
// @Failure 404 {object} map[string]string "News not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /news/{id} [put]
func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	// Получаем ID новости из URL
	vars := mux.Vars(r)
	newsIDParam := vars["id"]
	newsID, err := strconv.Atoi(newsIDParam)
	if err != nil {
		http.Error(w, `{"error": "Invalid news ID"}`, http.StatusBadRequest)
		return
	}

	// Декодируем тело запроса
	var req NewsUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Вызываем сервис обновления
	err = h.service.UpdateNewsService(newsID, req.Title, req.Content, req.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, `{"error": "News not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error": "Failed to update news"}`, http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated successfully"))
}

