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

	hub := domain.NewHub()
	go hub.Run()

	subscriptionUseCase := subscription.NewUseCase(l, rClient)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	v1.NewRouter(httpServer.App, cfg, l, subscriptionUseCase)

	// Start servers
	httpServer.Start()

	server := ws.NewHandler(hub, nil)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.HandleWebSocket(w, r)
	})

	email_message_handler := redis.EmailMessageHandler{Logger: l}
	emailSubscriber := redis.NewSubscriber("notification:email", email_message_handler, rClient, l)
	go emailSubscriber.Listen(context.Background())

	push_message_handler := redis.PushMessageHandler{Logger: l}
	pushSubscriber := redis.NewSubscriber("notification:push", push_message_handler, rClient, l)
	go pushSubscriber.Listen(context.Background())

	web_socket_message_handler := redis.WebSocketMessageHandler{Logger: l}
	webSocketSubscriber := redis.NewSubscriber("notification:web_socket", web_socket_message_handler, rClient, l)
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
