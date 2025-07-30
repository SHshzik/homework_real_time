package interfaces

import "context"

type SubscriptionRepository interface {
	RemoveSubscription(ctx context.Context, sType, userID string)
	AddSubscription(ctx context.Context, sType, userID string)
}
