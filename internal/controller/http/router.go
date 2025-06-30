package http

import (
	"net/http"

	"github.com/SHshzik/homework_real_time/config"
	"github.com/SHshzik/homework_real_time/internal/controller/http/middleware"
	v1 "github.com/SHshzik/homework_real_time/internal/controller/http/v1"
	"github.com/SHshzik/homework_real_time/internal/usecase"
	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// NewRouter -.
// Swagger spec:
// @title       V1 API
// @description Subscription CRUD
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *fiber.App, cfg *config.Config, l logger.Interface, t usecase.Subscription) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Swagger
	// if cfg.Swagger.Enabled {
	// 	app.Get("/swagger/*", swagger.HandlerDefault)
	// }

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewSubscriptionRoutes(apiV1Group, t, l)
	}
}
