package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rClient *redis.Client
}

func NewRepository(rClient *redis.Client) *Repository {
	return &Repository{
		rClient: rClient,
	}
}

func (r *Repository) AddSubscription(ctx context.Context, subscriptionType, subscriptionID string) {
	r.rClient.SAdd(ctx, subscriptionType, subscriptionID)
}

func (r *Repository) RemoveSubscription(ctx context.Context, subscriptionType, subscriptionID string) {
	r.rClient.SRem(ctx, subscriptionType, subscriptionID)
}
