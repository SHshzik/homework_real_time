package domain

type NotificationRepository interface {
	Save(notification *Notification) error
	GetByUserID(userID string) ([]*Notification, error)
	Delete(id string) error
}
