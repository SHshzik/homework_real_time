package redis

import (
	"github.com/SHshzik/homework_real_time/internal/domain"
)

type Repository struct {
	// TODO: Add Redis client
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Save(notification *domain.Notification) error {
	// TODO: Implement Redis save
	return nil
}

func (r *Repository) GetByUserID(userID string) ([]*domain.Notification, error) {
	// TODO: Implement Redis get
	return nil, nil
}

func (r *Repository) Delete(id string) error {
	// TODO: Implement Redis delete
	return nil
}
