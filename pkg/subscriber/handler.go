package subscriber

import "context"

//go:generate moq -stub -out mock/handler.go -pkg mock . Handler
type Handler interface {
	Call(ctx context.Context, message string) error
}
