package main

import (
	"log"

	"github.com/SHshzik/homework_real_time/config"
	"github.com/SHshzik/homework_real_time/internal/app"
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
