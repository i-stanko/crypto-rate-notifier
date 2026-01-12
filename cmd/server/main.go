package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/i-stanko/crypto-rate-notifier/internal/config"
	"github.com/i-stanko/crypto-rate-notifier/internal/handlers"
	"github.com/i-stanko/crypto-rate-notifier/internal/storage"
)

func main() {
	cfg := config.Load()

	store := storage.NewFileStore(cfg.SubscribersFile)
	api := handlers.NewAPIHandler(store)

	router := gin.Default()

	router.GET("/api/rate", api.GetBitcoinRate)
	router.POST("/api/subscribe", api.Subscribe)
	router.GET("/api/subscribers", api.ListSubscribers)

	addr := fmt.Sprintf(":%s", cfg.Port)
	router.Run(addr)
}
