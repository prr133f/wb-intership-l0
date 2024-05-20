package routes

import (
	"l0/internal/database"
	"l0/internal/domain"
	v1 "l0/internal/routes/v1"
	"l0/internal/view"
	"l0/pkg/cache"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func InitRouter(app *fiber.App, logger *zap.Logger, pg *database.Postgres, cache *cache.Cache) {
	db := database.NewDatabase(logger, pg)
	domain := domain.NewDomain(logger, db, cache)
	view := view.NewView(logger, domain)

	router := v1.NewRouter(app, view)
	router.InitRoutes()
	app.Get("/ping", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })
}
