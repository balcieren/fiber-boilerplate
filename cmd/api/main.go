package main

import (
	"github.com/balcieren/fiber-boilerplate/internal/api"
	v1 "github.com/balcieren/fiber-boilerplate/internal/api/v1"
	v2 "github.com/balcieren/fiber-boilerplate/internal/api/v2"
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/internal/database"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(database.NewPostgreSQL),
		fx.Provide(api.New),
		fx.Module("v1", v1.Module),
		fx.Module("v2", v2.Module),
		fx.Invoke(api.Start),
	).Run()
}
