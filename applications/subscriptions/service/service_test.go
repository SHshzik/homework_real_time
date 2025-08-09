package service

import (
	"context"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/domain"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/interfaces/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Subscribe(t *testing.T) {
	type fields struct {
		repository *mock.SubscriptionRepositoryMock
	}
	tests := []struct {
		name          string
		fields        fields
		repAddSubCall int
		subscription  *domain.Subscription
	}{
		{
			name: "Success",
			fields: fields{
				repository: &mock.SubscriptionRepositoryMock{},
			},
			repAddSubCall: 1,
			subscription:  &domain.Subscription{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				sRepository: tt.fields.repository,
			}
			err := s.Subscribe(context.Background(), tt.subscription)
			assert.NoError(t, err)
			if tt.fields.repository != nil {
				assert.Equal(t, tt.repAddSubCall, len(tt.fields.repository.AddSubscriptionCalls()))
			}
		})
	}
}

func TestService_Unsubscribe(t *testing.T) {
	type fields struct {
		repository *mock.SubscriptionRepositoryMock
	}
	tests := []struct {
		name          string
		fields        fields
		repRemSubCall int
		subscription  *domain.Subscription
	}{
		{
			name: "Success",
			fields: fields{
				repository: &mock.SubscriptionRepositoryMock{},
			},
			repRemSubCall: 1,
			subscription:  &domain.Subscription{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				sRepository: tt.fields.repository,
			}
			err := s.Unsubscribe(context.Background(), tt.subscription)
			assert.NoError(t, err)
			if tt.fields.repository != nil {
				assert.Equal(t, tt.repRemSubCall, len(tt.fields.repository.RemoveSubscriptionCalls()))
			}
		})
	}
}
