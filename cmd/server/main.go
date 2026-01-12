package main

import (
	"github.com/gin-gonic/gin"

	"crypto-rate-notifier/internal/handlers"
	"crypto-rate-notifier/internal/storage"
)

func main() {
	router := gin.Default()

	store := storage.NewFileStore("subscribers.txt")

	router.GET("/api/rate", handlers.GetCurrentBitcoinRate)
	router.POST("/api/subscribe", handlers.SubscribeEmail(store))
	router.POST("/api/sendEmails", handlers.ListSubscribers(store))

	router.Run(":8080")
}
