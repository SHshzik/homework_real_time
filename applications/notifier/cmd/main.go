package main

import (
	"context"
	"fmt"
	"github.com/SHshzik/homework_real_time/applications/notifier/adapters/redis"
	"github.com/SHshzik/homework_real_time/applications/notifier/config"
	"github.com/SHshzik/homework_real_time/applications/notifier/handlers"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/SHshzik/homework_real_time/pkg/subscriber"
	rds "github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig("configs/notifier.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	l := logger.New(cfg.Log.Level)

	redisOptions := &rds.Options{Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)}
	rClient := rds.NewClient(redisOptions)
	redisRepository := redis.NewRepository(rClient)

	emailMessageHandler := handlers.EmailMessageHandler{Logger: l, RedisRepository: redisRepository}
	emailSubscriber := subscriber.NewSubscriber("notification:email", emailMessageHandler, redisRepository, l)
	go emailSubscriber.Listen(context.Background())

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	s := <-interrupt
	l.Info("app - Run - signal: " + s.String())
}
