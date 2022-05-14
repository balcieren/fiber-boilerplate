package router

import (
	"github.com/balcieren/fiber-boilerplate/internal/api/handler"
	"github.com/balcieren/fiber-boilerplate/internal/api/middleware"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handler    *handler.Handler
	middleware *middleware.Middleware
	root, user fiber.Router
}

func New(app *fiber.App, handler *handler.Handler, middleware *middleware.Middleware) *Router {
	return &Router{
		root:       app.Group("/"),
		user:       app.Group("/users"),
		handler:    handler,
		middleware: middleware,
	}
}

func (r *Router) Setup() {
	r.root.Get("/", r.handler.Greet)

	r.user.Get("/", r.handler.GetAllUser)
	r.user.Post("/", r.handler.CreateUser)
}
