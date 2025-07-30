package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type subscriptionResponse struct {
	Message string `json:"message"`
}

func (h *HTTPServer) Subscribe(ctx *fiber.Ctx) error {
	formSubscription := subscriptionForm{}

	err := ctx.BodyParser(&formSubscription)
	if err != nil {
		h.l.Error(err, "http - v1 - subscribe - body parser")

		return errorResponse(ctx, http.StatusBadRequest, "Bad subscription params")
	}

	err = h.v.Struct(formSubscription)
	if err != nil {
		h.l.Error(err, "http - v1 - subscribe - validate")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	subscription := h.toSubscriptionDomain(formSubscription.SubType, formSubscription.UserID)

	err = h.s.Subscribe(ctx.UserContext(), subscription)
	if err != nil {
		h.l.Error(err, "http - v1 - subscribe - subscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "subscription not created")
	}

	return ctx.Status(http.StatusCreated).JSON(subscriptionResponse{
		Message: "Subscription created",
	})
}
