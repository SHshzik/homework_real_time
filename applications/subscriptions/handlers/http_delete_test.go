package handlers

import (
	"context"
	"fmt"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/domain"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/mock"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPServer_Unsubscribe(t *testing.T) {
	type fields struct {
		service *mock.ServiceMock
	}
	tests := []struct {
		name                   string
		fields                 fields
		reqBody                string
		serviceUnsubscribeCall int
		respStatus             int
	}{
		{
			name: "Success - unsubscribe user",
			fields: fields{
				service: &mock.ServiceMock{},
			},
			reqBody: `
				{
					"sub_type": "web_push",
					"user_id": "test_user"
				}
			`,
			serviceUnsubscribeCall: 1,
			respStatus:             204,
		},
		{
			name: "Fail - invalid body",
			reqBody: `
				{
					"sub_type": "web_push"
				}
			`,
			respStatus: 422,
		},
		{
			name: "Fail - service return error",
			fields: fields{
				service: &mock.ServiceMock{
					UnsubscribeFunc: func(ctx context.Context, subscription *domain.Subscription) error {
						return fmt.Errorf("some error")
					},
				},
			},
			reqBody: `
				{
					"sub_type": "web_push",
					"user_id": "test_user"
				}
			`,
			serviceUnsubscribeCall: 1,
			respStatus:             422,
		},
		{
			name: "Fail - invalid body",
			reqBody: `
				{
					"sub_type": "web_push"
				}
			`,
			respStatus: 422,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()

			h := &HTTPServer{
				s: tt.fields.service,
				l: logger.New("info"),
				v: validator.New(),
			}

			app.Delete("/subscriptions", h.Unsubscribe)

			req := httptest.NewRequest(http.MethodDelete, "/subscriptions", strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.respStatus, resp.StatusCode)

			_, err = io.ReadAll(resp.Body)
			assert.NoError(t, err)

			if tt.fields.service != nil {
				assert.Equal(t, tt.serviceUnsubscribeCall, len(tt.fields.service.UnsubscribeCalls()))
			}
		})
	}
}
