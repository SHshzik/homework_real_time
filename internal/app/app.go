package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/SHshzik/homework_real_time/config"
	ws "github.com/SHshzik/homework_real_time/internal/controller/websocket"
	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/domain/redis"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	rds "github.com/redis/go-redis/v9"
)

var addr = flag.String("addr", ":8081", "http service address")

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	hub := domain.NewHub()
	go hub.Run()

	server := ws.NewHandler(hub, nil)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.HandleWebSocket(w, r)
	})

	redisOptions := &rds.Options{Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)}
	client := rds.NewClient(redisOptions)

	email_message_handler := redis.EmailMessageHandler{Logger: l}
	emailSubscriber := redis.NewSubscriber("notification:email", email_message_handler, client, l)
	go emailSubscriber.Listen(context.Background())

	push_message_handler := redis.PushMessageHandler{Logger: l}
	pushSubscriber := redis.NewSubscriber("notification:push", push_message_handler, client, l)
	go pushSubscriber.Listen(context.Background())

	web_socket_message_handler := redis.WebSocketMessageHandler{Logger: l}
	webSocketSubscriber := redis.NewSubscriber("notification:web_socket", web_socket_message_handler, client, l)
	go webSocketSubscriber.Listen(context.Background())

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
