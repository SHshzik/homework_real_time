package subscriber

import (
	"context"
	rds "github.com/redis/go-redis/v9"
)

//go:generate moq -stub -out mock/repository.go -pkg mock . Repository
type Repository interface {
	Subscribe(ctx context.Context, sType string) *rds.PubSub
}
