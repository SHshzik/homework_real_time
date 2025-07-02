package v1

import (
	"net/http"

	"github.com/SHshzik/homework_real_time/internal/domain"
	"github.com/SHshzik/homework_real_time/internal/usecase"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type subscriptionRoutes struct {
	t usecase.Subscription
	l logger.Interface
	v *validator.Validate
}

func NewSubscriptionRoutes(apiV1Group fiber.Router, t usecase.Subscription, l logger.Interface) {
	r := &subscriptionRoutes{t, l, validator.New()}

	subscriptionGroup := apiV1Group.Group("/subscriptions")
	{
		subscriptionGroup.Post("/", r.subscribe)
		subscriptionGroup.Delete("/", r.unsubscribe)
	}
}

type subscriptionForm struct {
	SubType string `form:"sub_type"  validate:"required"`
	UserID  string `form:"user_id" validate:"required"`
}

// @Summary     Subscribe to a notification
// @Description Subscribe to a notification
// @ID          subscribe
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Success     201 {object} subscriptionResponse
// @Failure     422 {object} response
// @Failure     500 {object} response
// @Router      /subscriptions [post]
func (r *subscriptionRoutes) subscribe(ctx *fiber.Ctx) error {
	formSubscription := subscriptionForm{}

	err := ctx.BodyParser(&formSubscription)
	if err != nil {
		r.l.Error(err, "http - v1 - subscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "Bad subscription params")
	}

	err = r.v.Struct(formSubscription)
	if err != nil {
		r.l.Error(err, "http - v1 - subscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	subscription := domain.NewSubscription(formSubscription.SubType, formSubscription.UserID)

	err = r.t.Subscribe(ctx.UserContext(), subscription)
	if err != nil {
		r.l.Error(err, "http - v1 - subscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "subscription not created")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Subscription created",
	})
}

// @Summary     Unsubscribe from a notification
// @Description Unsubscribe from a notification
// @ID          unsubscribe
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Success     204
// @Failure     422 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /subscriptions [delete]
func (r *subscriptionRoutes) unsubscribe(ctx *fiber.Ctx) error {
	formSubscription := subscriptionForm{}

	err := ctx.BodyParser(&formSubscription)
	if err != nil {
		r.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "Bad subscription params")
	}

	err = r.v.Struct(formSubscription)
	if err != nil {
		r.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	subscription := domain.NewSubscription(formSubscription.SubType, formSubscription.UserID)

	err = r.t.Unsubscribe(ctx.UserContext(), subscription)
	if err != nil {
		r.l.Error(err, "http - v1 - unsubscribe")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "subscription not deleted")
	}

	return ctx.Status(http.StatusNoContent).JSON(fiber.Map{
		"message": "Subscription deleted",
	})
}
