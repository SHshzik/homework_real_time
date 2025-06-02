package domain

type Notification struct {
	ID        string
	UserID    string
	Type      string
	Content   string
	CreatedAt int64
}
