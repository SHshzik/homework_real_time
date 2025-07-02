package subscription

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/domain/redis"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type UseCase struct {
	l               *logger.Logger
	RedisRepository *redis.Repository
}

func NewUseCase(l *logger.Logger, redisRepository *redis.Repository) *UseCase {
	return &UseCase{l: l, RedisRepository: redisRepository}
}

func (uc *UseCase) Subscribe(ctx context.Context, subscription *domain.Subscription) error {
	uc.RedisRepository.AddSubscription(ctx, subscription.Type, subscription.UserID)

	return nil
}

func (uc *UseCase) Unsubscribe(ctx context.Context, subscription *domain.Subscription) error {
	uc.RedisRepository.RemoveSubscription(ctx, subscription.Type, subscription.UserID)

	return nil
}
