package redis

import (
	"context"
	"encoding/json"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type PushMessageHandler struct {
	Logger          *logger.Logger
	RedisRepository *Repository
}

func (h PushMessageHandler) Call(ctx context.Context, message string) error {
	messageEntity := new(domain.Message)

	err := json.Unmarshal([]byte(message), messageEntity)
	if err != nil {
		return err
	}

	subscriptions := h.RedisRepository.FetchSubscriptions(ctx, domain.SubscriptionTypePush)
	for _, subscription := range subscriptions {
		h.Logger.Info("send push message (%#v) to subscription: %#v", messageEntity, subscription)
	}

	return nil
}
