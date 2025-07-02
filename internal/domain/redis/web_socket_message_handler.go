package redis

import (
	"context"
	"encoding/json"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type WebSocketMessageHandler struct {
	Logger  *logger.Logger
	RClient *redis.Client
}

func (h WebSocketMessageHandler) Call(ctx context.Context, message string) error {
	messageEntity := new(domain.Message)
	json.Unmarshal([]byte(message), messageEntity)

	subscriptions := h.RClient.SMembers(ctx, domain.SubscriptionTypeWS)
	for _, subscription := range subscriptions.Val() {
		h.Logger.Info("send web socket message (%#v) to subscription: %#v", messageEntity, subscription)
	}

	return nil
}
