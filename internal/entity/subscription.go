package entity

type Subscription struct {
	Type   string
	UserID string
}

func NewSubscription(subType, userID string) *Subscription {
	return &Subscription{
		Type:   subType,
		UserID: userID,
	}
}
