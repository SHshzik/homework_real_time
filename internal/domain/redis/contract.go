package redis

import "context"

type Handler interface {
	Call(ctx context.Context, message string) error
}
