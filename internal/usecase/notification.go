package usecase

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/domain"
)

type NotificationUseCase struct{}

func NewNotificationUseCase() *NotificationUseCase {
	return &NotificationUseCase{}
}

func (uc *NotificationUseCase) SendNotification(ctx context.Context, notification *domain.Notification) error {
	return nil
}
