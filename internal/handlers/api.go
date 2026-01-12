package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/i-stanko/crypto-rate-notifier/internal/storage"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	store storage.SubscriberStore
}

func NewAPIHandler(store storage.SubscriberStore) *APIHandler {
	return &APIHandler{store: store}
}

// GET /api/rate
func (h *APIHandler) GetBitcoinRate(c *gin.Context) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch rate"})
		return
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid response"})
		return
	}

	rate := result["bitcoin"]["uah"]
	c.JSON(http.StatusOK, gin.H{"rate": rate})
}

// POST /api/subscribe
func (h *APIHandler) Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if err := h.store.Add(email); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email subscribed"})
}

// POST /api/sendEmails
func (h *APIHandler) ListSubscribers(c *gin.Context) {
	subscribers, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read subscribers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscribers": subscribers})
}
