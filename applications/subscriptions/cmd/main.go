package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SHshzik/homework_real_time/applications/subscriptions/adapters/redis"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/config"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/handlers"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/service"
	"github.com/SHshzik/homework_real_time/pkg/httpserver"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/gofiber/fiber/v2"
	rds "github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.NewConfig("configs/subscriptions.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	l := logger.New(cfg.Log.Level)

	redisOptions := &rds.Options{Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)}
	rClient := rds.NewClient(redisOptions)
	redisRepository := redis.NewRepository(rClient)

	subService := service.NewService(redisRepository)

	server := httpserver.New(httpserver.Port(cfg.HTTP.Port))

	{
		server.App.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })
	}

	{
		subServer := handlers.NewHTTPServer(subService, l)
		v1Ground := server.App.Group("/v1")
		v1Ground.Post("/subscriptions", subServer.Subscribe)
		v1Ground.Delete("/subscriptions", subServer.Unsubscribe)
	}

	server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
