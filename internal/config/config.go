package config

import (
	"log"
	"os"
)

type Config struct {
	Port            string
	SubscribersFile string
}

func Load() *Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	file := os.Getenv("SUBSCRIBERS_FILE")
	if file == "" {
		file = "subscribers.txt"
	}

	log.Printf("config loaded: port=%s subscribers_file=%s", port, file)

	return &Config{
		Port:            port,
		SubscribersFile: file,
	}
}
