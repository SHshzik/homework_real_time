package domain

const (
	SubscriptionTypeEmail = "email"
	SubscriptionTypePush  = "push"
	SubscriptionTypeWS    = "ws"
)

type Subscription struct {
	Type   string
	UserID string
}
