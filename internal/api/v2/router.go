package v2

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/v2/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	user    fiber.Router
	handler *handler.Handler
}

func NewRouter(app *fiber.App, handlerV2 *handler.Handler) *Router {
	return &Router{
		user:    app.Group("/v2/users"),
		handler: handlerV2,
	}
}

func (r *Router) Setup() {
	r.user.Get("/", r.handler.GetAllUser)
	r.user.Post("/", r.handler.CreateUser)
}
