package main

import (
	"github.com/balcieren/fiber-boilerplate/internal/api"
	"github.com/balcieren/fiber-boilerplate/internal/api/handler"
	"github.com/balcieren/fiber-boilerplate/internal/api/middleware"
	"github.com/balcieren/fiber-boilerplate/internal/api/router"
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/internal/database"
	"github.com/balcieren/fiber-boilerplate/pkg/service"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(database.NewPostgreSQL),
		fx.Provide(api.New),
		fx.Provide(service.New),
		fx.Provide(handler.New),
		fx.Provide(middleware.New),
		fx.Provide(router.New),
		fx.Invoke(api.Start),
	).Run()
}
