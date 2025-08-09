package interfaces

import "context"

//go:generate moq -stub -out mock/subscription_repository.go -pkg mock . SubscriptionRepository
type SubscriptionRepository interface {
	RemoveSubscription(ctx context.Context, sType, userID string)
	AddSubscription(ctx context.Context, sType, userID string)
}
