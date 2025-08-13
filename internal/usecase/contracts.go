package usecase

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/domain"
)

type (
	// User -.
	Subscription interface {
		Subscribe(ctx context.Context, subscription *domain.Subscription) error
		Unsubscribe(ctx context.Context, subscription *domain.Subscription) error
	}
)
