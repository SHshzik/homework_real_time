package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SHshzik/homework_real_time/applications/notifier/domain"
	"github.com/SHshzik/homework_real_time/applications/notifier/interfaces"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type EmailMessageHandler struct {
	Logger          *logger.Logger
	RedisRepository interfaces.SubscriptionRepository
}

func (h EmailMessageHandler) Call(ctx context.Context, message string) error {
	messageEntity := new(domain.Message)

	err := json.Unmarshal([]byte(message), messageEntity)
	if err != nil {
		return err
	}
	fmt.Println(messageEntity)

	subscriptions := h.RedisRepository.FetchSubscriptions(ctx, domain.SubscriptionTypeEmail)
	for _, subscription := range subscriptions {
		h.Logger.Info("send message (%#v) to subscription: %#v", messageEntity, subscription)
	}

	return nil
}
