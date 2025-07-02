package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SHshzik/homework_real_time/config"
	v1 "github.com/SHshzik/homework_real_time/internal/controller/http"
	ws "github.com/SHshzik/homework_real_time/internal/controller/websocket"
	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/domain/redis"
	"github.com/SHshzik/homework_real_time/internal/usecase/subscription"
	"github.com/SHshzik/homework_real_time/pkg/httpserver"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	rds "github.com/redis/go-redis/v9"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	redisOptions := &rds.Options{Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)}
	rClient := rds.NewClient(redisOptions)
	redisRepository := redis.NewRepository(rClient)

	hub := domain.NewHub()
	go hub.Run()

	subscriptionUseCase := subscription.NewUseCase(l, redisRepository)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	v1.NewRouter(httpServer.App, cfg, l, subscriptionUseCase)

	// Start servers
	httpServer.Start()

	server := ws.NewHandler(hub, nil, l)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := server.HandleWebSocket(w, r)
		if err != nil {
			l.Error(fmt.Errorf("app - Run - server.HandleWebSocket: %w", err))
		}
	})

	emailMessageHandler := redis.EmailMessageHandler{Logger: l, RedisRepository: redisRepository}
	emailSubscriber := redis.NewSubscriber("notification:email", emailMessageHandler, redisRepository, l)

	go emailSubscriber.Listen(context.Background())

	pushMessageHandler := redis.PushMessageHandler{Logger: l, RedisRepository: redisRepository}
	pushSubscriber := redis.NewSubscriber("notification:push", pushMessageHandler, redisRepository, l)

	go pushSubscriber.Listen(context.Background())

	webSocketMessageHandler := redis.WebSocketMessageHandler{Logger: l, RedisRepository: redisRepository}
	webSocketSubscriber := redis.NewSubscriber("notification:web_socket", webSocketMessageHandler, redisRepository, l)

	go webSocketSubscriber.Listen(context.Background())

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
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
