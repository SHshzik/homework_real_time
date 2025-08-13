package handlers

import (
	"context"
	"github.com/SHshzik/homework_real_time/applications/notifier/domain"
	"github.com/SHshzik/homework_real_time/applications/notifier/interfaces/mock"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushMessageHandler_Call(t *testing.T) {
	type fields struct {
		rep *mock.SubscriptionRepositoryMock
	}
	tests := []struct {
		name           string
		fields         fields
		message        string
		wantErr        assert.ErrorAssertionFunc
		wantErrMessage string
	}{
		{
			name: "Success - message successful send",
			fields: fields{
				rep: &mock.SubscriptionRepositoryMock{
					FetchSubscriptionsFunc: func(ctx context.Context, subscriptionType domain.SubscriptionType) []string {
						return []string{}
					},
				},
			},
			message: "{}",
			wantErr: assert.NoError,
		},
		{
			name:    "Failure - invalid message",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := PushMessageHandler{
				l:   logger.New("info"),
				rep: tt.fields.rep,
			}

			err := h.Call(context.Background(), tt.message)
			tt.wantErr(t, err, tt.wantErrMessage)
		})
	}
}
