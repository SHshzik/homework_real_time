package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPServer) Unsubscribe(ctx *fiber.Ctx) error {
	formSubscription := subscriptionForm{}

	err := ctx.BodyParser(&formSubscription)
	if err != nil {
		h.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "Bad subscription params")
	}

	err = h.v.Struct(formSubscription)
	if err != nil {
		h.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	subscription := h.toSubscriptionDomain(formSubscription.SubType, formSubscription.UserID)

	err = h.s.Unsubscribe(ctx.UserContext(), subscription)
	if err != nil {
		h.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "subscription not deleted")
	}

	return ctx.Status(http.StatusNoContent).JSON(subscriptionResponse{
		Message: "Subscription deleted",
	})
}
