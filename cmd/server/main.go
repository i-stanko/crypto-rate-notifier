package main

import (
	"github.com/gin-gonic/gin"

	"github.com/i-stanko/crypto-rate-notifier/internal/handlers"
	"github.com/i-stanko/crypto-rate-notifier/internal/storage"
)

func main() {
	store := storage.NewFileStore("subscribers.txt")
	api := handlers.NewAPIHandler(store)

	router := gin.Default()

	router.GET("/api/rate", api.GetBitcoinRate)
	router.POST("/api/subscribe", api.Subscribe)
	router.GET("/api/subscribers", api.ListSubscribers)

	router.Run(":8080")
}
