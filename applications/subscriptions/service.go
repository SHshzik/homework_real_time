package subscriptions

import (
	"context"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/domain"
)

//go:generate moq -stub -out mock/service.go -pkg mock . Service
type Service interface {
	Subscribe(ctx context.Context, subscription *domain.Subscription) error
	Unsubscribe(ctx context.Context, subscription *domain.Subscription) error
}
