package handlers

import (
	"github.com/SHshzik/homework_real_time/applications/subscriptions"
	"github.com/SHshzik/homework_real_time/applications/subscriptions/domain"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type subscriptionForm struct {
	SubType string `json:"sub_type" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}

type HTTPServer struct {
	s subscriptions.Service
	l logger.Interface
	v *validator.Validate
}

func NewHTTPServer(s subscriptions.Service, l logger.Interface) *HTTPServer {
	return &HTTPServer{s: s, l: l, v: validator.New()}
}

func (h *HTTPServer) toSubscriptionDomain(subType, userID string) *domain.Subscription {
	return &domain.Subscription{
		Type:   subType,
		UserID: userID,
	}
}
