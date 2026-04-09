package main

import (
	"fmt"
	"log"
	"tourism-backend/config"
	"tourism-backend/storage"
)

func main() {
	cfg := config.Load()

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	fmt.Println("Connected to database")
	fmt.Println("Server starting on port", cfg.ServerPort)
}
