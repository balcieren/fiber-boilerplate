package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Greet(c *fiber.Ctx) error {
	return c.JSON("The best fiber-boilerplate")
}
