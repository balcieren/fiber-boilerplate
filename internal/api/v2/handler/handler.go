package handler

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v2/service"
	"github.com/balcieren/fiber-boilerplate/internal/config"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		service,
		config,
	}
}
