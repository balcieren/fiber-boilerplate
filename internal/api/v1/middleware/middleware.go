package middleware

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v1/service"
	"github.com/balcieren/fiber-boilerplate/internal/config"
)

type Middleware struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, config *config.Config) *Middleware {
	return &Middleware{
		service,
		config,
	}
}
