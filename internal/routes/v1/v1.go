package v1

import (
	"l0/internal/view"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App  *fiber.App
	View *view.View
}

func NewRouter(app *fiber.App, view *view.View) *Router {
	return &Router{
		App:  app,
		View: view,
	}
}

type Route struct {
	Group *fiber.Router
	View  *view.View
}

func (r *Router) InitRoutes() {
	v1 := r.App.Group("/v1")
	route := &Route{
		Group: &v1,
		View:  r.View,
	}
	route.Data()
}
