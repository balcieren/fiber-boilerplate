package service

import (
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/pkg/ent"
)

type Service struct {
	db     *ent.Client
	config *config.Config
}

func New(db *ent.Client, config *config.Config) *Service {
	return &Service{
		db:     db,
		config: config,
	}
}
