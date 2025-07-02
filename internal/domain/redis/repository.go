package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rClient *redis.Client
}

func NewRepository(rClient *redis.Client) *Repository {
	return &Repository{rClient: rClient}
}

func (r *Repository) FetchSubscriptions(ctx context.Context, subscriptionType string) []string {
	return r.rClient.SMembers(ctx, subscriptionType).Val()
}

func (r *Repository) Subscribe(ctx context.Context, subscriptionType string) *redis.PubSub {
	return r.rClient.Subscribe(ctx, subscriptionType)
}

func (r *Repository) AddSubscription(ctx context.Context, subscriptionType string, subscriptionID string) {
	r.rClient.SAdd(ctx, subscriptionType, subscriptionID)
}

func (r *Repository) RemoveSubscription(ctx context.Context, subscriptionType string, subscriptionID string) {
	r.rClient.SRem(ctx, subscriptionType, subscriptionID)
}
