package subscription

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/entity"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	l       *logger.Logger
	rClient *redis.Client
}

func NewUseCase(l *logger.Logger, rClient *redis.Client) *UseCase {
	return &UseCase{l: l, rClient: rClient}
}

func (uc *UseCase) Subscribe(ctx context.Context, subscription *entity.Subscription) error {
	uc.rClient.SAdd(ctx, subscription.Type, subscription.UserID)

	return nil
}

func (uc *UseCase) Unsubscribe(ctx context.Context, subscription *entity.Subscription) error {
	uc.rClient.SRem(ctx, subscription.Type, subscription.UserID)

	return nil
}
