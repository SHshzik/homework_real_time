package handlers

import (
	"context"
	"encoding/json"
	"github.com/SHshzik/homework_real_time/applications/notifier/domain"
	"github.com/SHshzik/homework_real_time/applications/notifier/interfaces"
	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type PushMessageHandler struct {
	l   logger.Interface
	rep interfaces.SubscriptionRepository
}

func NewPushMessageHandler(l logger.Interface, rep interfaces.SubscriptionRepository) *PushMessageHandler {
	return &PushMessageHandler{
		l:   l,
		rep: rep,
	}
}

func (h PushMessageHandler) Call(ctx context.Context, message string) error {
	messageEntity := new(domain.Message)

	err := json.Unmarshal([]byte(message), messageEntity)
	if err != nil {
		return err
	}

	subscriptions := h.rep.FetchSubscriptions(ctx, domain.SubscriptionTypePush)
	for _, subscription := range subscriptions {
		h.l.Info("send push message (%#v) to subscription: %#v", messageEntity, subscription)
	}

	return nil
}
