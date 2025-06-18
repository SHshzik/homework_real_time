package redis

import (
	"context"

	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type PushMessageHandler struct {
	Logger *logger.Logger
}

func (h PushMessageHandler) Call(ctx context.Context, message string) error {
	h.Logger.Info("push message received:", message)

	return nil
}
