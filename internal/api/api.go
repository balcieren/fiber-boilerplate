package api

import (
	"context"
	"fmt"

	"github.com/balcieren/fiber-boilerplate/internal/api/helper"
	v1 "github.com/balcieren/fiber-boilerplate/internal/api/v1"
	v2 "github.com/balcieren/fiber-boilerplate/internal/api/v2"

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

func Start(lifecycle fx.Lifecycle, app *fiber.App, config *config.Config, rv1 *v1.Router, rv2 *v2.Router) {
	app.Use(logger.New())

	rv1.Setup()
	rv2.Setup()

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
