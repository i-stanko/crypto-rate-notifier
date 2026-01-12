package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/i-stanko/crypto-rate-notifier/internal/config"
	"github.com/i-stanko/crypto-rate-notifier/internal/handlers"
	"github.com/i-stanko/crypto-rate-notifier/internal/storage"
)

func main() {
	cfg := config.Load()

	store := storage.NewFileStore(cfg.SubscribersFile)
	api := handlers.NewAPIHandler(store)
	health := handlers.NewHealthHandler(store)

	router := gin.Default()
	router.GET("/api/rate", api.GetBitcoinRate)
	router.POST("/api/subscribe", api.Subscribe)
	router.GET("/api/subscribers", api.ListSubscribers)
	router.GET("/healthz", health.Healthz)
	router.GET("/readyz", health.Readyz)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	// run server
	go func() {
		log.Printf("server started on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server exited gracefully")
}
