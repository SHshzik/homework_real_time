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

func TestHTTPServer_Subscribe(t *testing.T) {
	type fields struct {
		service *mock.ServiceMock
	}
	tests := []struct {
		name                 string
		fields               fields
		reqBody              string
		expectedRespBody     string
		serviceSubscribeCall int
		respStatus           int
	}{
		{
			name: "Success - subscribe user",
			fields: fields{
				service: &mock.ServiceMock{},
			},
			reqBody: `
				{
					"sub_type": "web_push",
					"user_id": "test_user"
				}
			`,
			expectedRespBody:     `{"message": "Subscription created"}`,
			serviceSubscribeCall: 1,
			respStatus:           201,
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
					SubscribeFunc: func(ctx context.Context, subscription *domain.Subscription) error {
						return fmt.Errorf("Some error")
					},
				},
			},
			reqBody: `
				{
					"sub_type": "web_push",
					"user_id": "test_user"
				}
			`,
			serviceSubscribeCall: 1,
			respStatus:           422,
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

			app.Post("/subscriptions", h.Subscribe)

			req := httptest.NewRequest(http.MethodPost, "/subscriptions", strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.respStatus, resp.StatusCode)

			b, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			expectedBody := tt.expectedRespBody
			if expectedBody != "" {
				assert.JSONEq(t, expectedBody, string(b))
			}

			if tt.fields.service != nil {
				assert.Equal(t, tt.serviceSubscribeCall, len(tt.fields.service.SubscribeCalls()))
			}
		})
	}
}
