package websocket

import (
	"github.com/SHshzik/homework_real_time/internal/usecase"
)

type Handler struct {
	notificationUseCase *usecase.NotificationUseCase
}

func NewHandler(notificationUseCase *usecase.NotificationUseCase) *Handler {
	return &Handler{
		notificationUseCase: notificationUseCase,
	}
}

func (h *Handler) HandleWebSocket() error {
	// TODO: Implement WebSocket handling
	return nil
}
