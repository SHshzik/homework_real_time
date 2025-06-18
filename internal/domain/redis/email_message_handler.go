package redis

import (
	"context"

	"github.com/SHshzik/homework_real_time/pkg/logger"
)

type EmailMessageHandler struct {
	Logger *logger.Logger
}

func (h EmailMessageHandler) Call(ctx context.Context, message string) error {
	h.Logger.Info("email message received:", message)

	return nil
}
