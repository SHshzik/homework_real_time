package usecase

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/domain"
)

type NotificationUseCase struct {
	repo domain.NotificationRepository
}

func NewNotificationUseCase(repo domain.NotificationRepository) *NotificationUseCase {
	return &NotificationUseCase{
		repo: repo,
	}
}

func (uc *NotificationUseCase) SendNotification(ctx context.Context, notification *domain.Notification) error {
	return uc.repo.Save(notification)
}
