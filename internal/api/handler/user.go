package handler

import (
	"context"

	"github.com/balcieren/fiber-boilerplate/pkg/ent"
	"github.com/balcieren/fiber-boilerplate/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetAllUser(c *fiber.Ctx) error {
	query := new(service.GetAllUserRequest)
	if err := c.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	users, err := h.service.GetAllUser(context.Background(), service.GetAllUserRequest{
		PerPage: query.PerPage,
	})
	if ent.IsNotFound(err) {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(users)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	body := new(service.CreateUserRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateUser(context.Background(), service.CreateUserRequest{
		Name: body.Name,
		Age:  body.Age,
	}); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("User has created")
}
