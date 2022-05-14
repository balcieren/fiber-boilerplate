package service

import (
	"context"

	"github.com/balcieren/fiber-boilerplate/pkg/ent"
)

type GetAllUserRequest struct {
	PerPage int `json:"perPage"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *Service) GetAllUser(ctx context.Context, req GetAllUserRequest) ([]*ent.User, error) {
	return s.db.User.Query().Limit(req.PerPage).All(ctx)
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) error {
	return s.db.User.Create().SetName(req.Name).SetAge(req.Age).Exec(ctx)
}
