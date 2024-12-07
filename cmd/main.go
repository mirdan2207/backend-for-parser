package main

import (
	"log"
	"testing_po/config"
	"testing_po/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	application := app.New(cfg)

	log.Println("Server is running on port:", cfg.ServerPort)
	if err := application.Run(); err != nil {
        log.Fatalf("Error running the server: %v", err)
    }
}