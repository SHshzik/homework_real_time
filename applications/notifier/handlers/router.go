package handlers

import (
	"net/http"

	"github.com/SHshzik/homework_real_time/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, l logger.Interface) {
	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })
}
