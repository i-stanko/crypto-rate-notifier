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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch rate",
		})
		c.Writer.WriteString("\n")
		return
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "invalid response",
		})
		c.Writer.WriteString("\n")
		return
	}

	rate := result["bitcoin"]["uah"]
	c.IndentedJSON(http.StatusOK, gin.H{
		"rate": rate,
	})
	c.Writer.WriteString("\n")
}

// POST /api/subscribe
func (h *APIHandler) Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "email is required",
		})
		c.Writer.WriteString("\n")
		return
	}

	if err := h.store.Add(email); err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		c.Writer.WriteString("\n")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "email subscribed",
	})
	c.Writer.WriteString("\n")
}

// GET /api/subscribers
func (h *APIHandler) ListSubscribers(c *gin.Context) {
	subscribers, err := h.store.List()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read subscribers",
		})
		c.Writer.WriteString("\n")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"subscribers": subscribers,
	})
	c.Writer.WriteString("\n")
}
