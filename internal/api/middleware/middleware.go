package middleware

import (
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/pkg/service"
)

type Middleware struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, config *config.Config) *Middleware {
	return &Middleware{
		service: service,
		config:  config,
	}
}
