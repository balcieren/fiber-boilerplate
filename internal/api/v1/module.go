package v1

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v1/handler"
	"github.com/balcieren/fiber-boilerplate/internal/api/v1/middleware"
	"github.com/balcieren/fiber-boilerplate/internal/api/v1/service"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	middleware.New,
	NewRouter,
)
