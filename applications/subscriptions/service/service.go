package service

import (
	"context"

	"github.com/SHshzik/homework_real_time/applications/subscriptions/domain"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/interfaces"
)

type Service struct {
	sRepository interfaces.SubscriptionRepository
}

func NewService(sRepository interfaces.SubscriptionRepository) *Service {
	return &Service{
		sRepository: sRepository,
	}
}

func (s *Service) Subscribe(ctx context.Context, subscription *domain.Subscription) error {
	s.sRepository.AddSubscription(ctx, subscription.Type, subscription.UserID)

	return nil
}

func (s *Service) Unsubscribe(ctx context.Context, subscription *domain.Subscription) error {
	s.sRepository.RemoveSubscription(ctx, subscription.Type, subscription.UserID)

	return nil
}
