package v2

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v2/handler"
	"github.com/balcieren/fiber-boilerplate/internal/api/v2/middleware"
	"github.com/balcieren/fiber-boilerplate/internal/api/v2/service"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	middleware.New,
	NewRouter,
)
