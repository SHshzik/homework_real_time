package redis

import (
	"context"
	"encoding/json"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type WebSocketMessageHandler struct {
	Logger          *logger.Logger
	RedisRepository *Repository
}

func (h WebSocketMessageHandler) Call(ctx context.Context, message string) error {
	messageEntity := new(domain.Message)
	json.Unmarshal([]byte(message), messageEntity)

	subscriptions := h.RedisRepository.FetchSubscriptions(ctx, domain.SubscriptionTypeWS)
	for _, subscription := range subscriptions {
		h.Logger.Info("send web socket message (%#v) to subscription: %#v", messageEntity, subscription)
	}

	return nil
}
