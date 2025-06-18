package redis

import (
	"context"

	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type WebSocketMessageHandler struct {
	Logger *logger.Logger
}

func (h WebSocketMessageHandler) Call(ctx context.Context, message string) error {
	h.Logger.Info("web socket message received:", message)

	return nil
}
