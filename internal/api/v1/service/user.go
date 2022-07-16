package service

import (
	"context"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAllUserRequest struct {
	PerPage int `json:"perPage"`
}

type GetAllUserResponse = []User

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (s *Service) GetAllUser(ctx context.Context, req GetAllUserRequest) (*GetAllUserResponse, error) {
	var users []User
	u, err := s.db.User.Query().Limit(req.PerPage).All(ctx)

	for _, user := range u {
		users = append(users, User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return &users, err
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) error {
	return s.db.User.Create().SetName(req.Name).Exec(ctx)
}
