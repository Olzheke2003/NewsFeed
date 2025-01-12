package transport

import (
	"encoding/json"
	"net/http"

	"github.com/Olzheke2003/NewsFeed/internal/services"
)

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
