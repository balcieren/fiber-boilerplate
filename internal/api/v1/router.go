package v1

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v1/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	user    fiber.Router
	handler *handler.Handler
}

func NewRouter(app *fiber.App, handlerV1 *handler.Handler) *Router {
	return &Router{
		user:    app.Group("/v1/users"),
		handler: handlerV1,
	}
}

func (r *Router) Setup() {
	r.user.Get("/", r.handler.GetAllUser)
	r.user.Post("/", r.handler.CreateUser)
}
