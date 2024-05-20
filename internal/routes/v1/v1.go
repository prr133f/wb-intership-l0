package v1

import (
	"l0/internal/view"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App  *fiber.App
	View view.IFace
}

func NewRouter(app *fiber.App, view view.IFace) *Router {
	return &Router{
		App:  app,
		View: view,
	}
}

type Route struct {
	Group *fiber.Router
	View  view.IFace
}

func (r *Router) InitRoutes() {
	v1 := r.App.Group("/v1")
	route := &Route{
		Group: &v1,
		View:  r.View,
	}
	route.Data()
}
