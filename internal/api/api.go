package api

import (
	"context"
	"fmt"

	"github.com/balcieren/fiber-boilerplate/internal/api/helper"
	"github.com/balcieren/fiber-boilerplate/internal/api/router"
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

func New() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:      "Fiber-Boilerplate",
		ErrorHandler: helper.ErrorHandler,
	})
}

func Start(lifecycle fx.Lifecycle, app *fiber.App, config *config.Config, router *router.Router) {
	app.Use(logger.New())

	router.Setup()

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(fmt.Sprintf(":%s", config.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
