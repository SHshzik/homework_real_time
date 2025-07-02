package redis

import (
	"context"
	"fmt"

	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type Subscriber struct {
	Name       string
	Handler    Handler
	repository *Repository
	Logger     *logger.Logger
}

func NewSubscriber(name string, handler Handler, repository *Repository, logger *logger.Logger) *Subscriber {
	return &Subscriber{
		Name:       name,
		Handler:    handler,
		repository: repository,
		Logger:     logger,
	}
}

func (s *Subscriber) Listen(ctx context.Context) error {
	pubsub := s.repository.Subscribe(ctx, s.Name)
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
