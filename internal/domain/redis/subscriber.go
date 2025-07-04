package redis

import (
	"context"

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

func (s *Subscriber) Listen(ctx context.Context) {
	pubsub := s.repository.Subscribe(ctx, s.Name)
	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			if err := s.Handler.Call(ctx, msg.Payload); err != nil {
				s.Logger.Error("handler error: %v\n", err)
			}
		case <-ctx.Done():
			s.Logger.Info("context done")
		}
	}
}
