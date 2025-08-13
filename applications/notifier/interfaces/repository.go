package interfaces

import (
	"context"
	"github.com/SHshzik/homework_real_time/applications/notifier/domain"
)

//go:generate moq -stub -out mock/subscription_repository.go -pkg mock . SubscriptionRepository
type SubscriptionRepository interface {
	FetchSubscriptions(ctx context.Context, subscriptionType domain.SubscriptionType) []string
}
