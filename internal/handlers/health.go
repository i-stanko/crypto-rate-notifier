package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-stanko/crypto-rate-notifier/internal/storage"
)

type HealthHandler struct {
	store storage.SubscriberStore
}

func NewHealthHandler(store storage.SubscriberStore) *HealthHandler {
	return &HealthHandler{store: store}
}

// GET /healthz
func (h *HealthHandler) Healthz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	c.Writer.WriteString("\n")
}

// GET /readyz
func (h *HealthHandler) Readyz(c *gin.Context) {
	_, err := h.store.List()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{
			"status": "not ready",
		})
		c.Writer.WriteString("\n")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "ready",
	})
	c.Writer.WriteString("\n")
}
