package main

import (
	"log"

	"github.com/evrone/go-clean-template/internal/app"
	"homework_crud/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
