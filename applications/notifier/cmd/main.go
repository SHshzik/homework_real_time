package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SHshzik/homework_real_time/applications/notifier/config"
	"github.com/SHshzik/homework_real_time/applications/notifier/handlers"
	"github.com/SHshzik/homework_real_time/pkg/httpserver"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

func main() {
	cfg, err := config.NewConfig("configs/notifier.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	l := logger.New(cfg.Log.Level)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))

	handlers.NewRouter(httpServer.App, l)

	// Start servers
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
