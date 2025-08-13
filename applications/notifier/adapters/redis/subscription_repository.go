package redis

import (
	"context"
	"github.com/SHshzik/homework_real_time/applications/notifier/domain"
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

func (r *Repository) Subscribe(ctx context.Context, subscriptionType string) *redis.PubSub {
	return r.rClient.Subscribe(ctx, subscriptionType)
}

func (r *Repository) FetchSubscriptions(ctx context.Context, subscriptionType domain.SubscriptionType) []string {
	return r.rClient.SMembers(ctx, string(subscriptionType)).Val()
}
