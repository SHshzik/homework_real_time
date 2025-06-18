package redis

import (
	"context"
	"fmt"

	"github.com/SHshzik/homework_real_time/pkg/logger"
	rds "github.com/redis/go-redis/v9"
)

type Subscriber struct {
	Name    string
	Handler Handler
	client  *rds.Client
	Logger  *logger.Logger
}

func NewSubscriber(name string, handler Handler, client *rds.Client, logger *logger.Logger) *Subscriber {
	return &Subscriber{
		Name:    name,
		Handler: handler,
		client:  client,
		Logger:  logger,
	}
}

func (s *Subscriber) Listen(ctx context.Context) error {
	pubsub := s.client.Subscribe(ctx, s.Name)
	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			if err := s.Handler.Call(ctx, msg.Payload); err != nil {
				fmt.Printf("handler error: %v\n", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Пример использования:
// type EmailMessageHandler struct{}
// func (h EmailMessageHandler) Call(ctx context.Context, message string) error {
//     // Здесь логика отправки email
//     return nil
// }
//
// client := rds.NewClient(&rds.Options{Addr: "localhost:6379"})
// emailSubscriber := NewSubscriber("notification:email", EmailMessageHandler{}, client)
// go emailSubscriber.Listen(context.Background())
