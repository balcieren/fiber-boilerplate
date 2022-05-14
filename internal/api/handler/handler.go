package handler

import (
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/pkg/service"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  config,
	}
}
