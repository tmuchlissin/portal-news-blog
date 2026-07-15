package app

import (
	"log"
	"portal-news-blog/config"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
}
