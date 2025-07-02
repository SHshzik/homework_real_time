package usecase

import (
	"context"

	"github.com/SHshzik/homework_real_time/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test
type (
	// User -.
	Subscription interface {
		Subscribe(ctx context.Context, subscription *domain.Subscription) error
		Unsubscribe(ctx context.Context, subscription *domain.Subscription) error
	}
)
