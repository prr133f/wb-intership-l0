package routes

import (
	"l0/internal/database"
	"l0/internal/domain"
	v1 "l0/internal/routes/v1" //nolint
	"l0/internal/view"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func InitRouter(app *fiber.App, logger *zap.Logger, pg *database.Postgres) {
	db := database.NewDatabase(logger, pg)
	domain := domain.NewDomain(logger, db)
	view := view.NewView(logger, domain)

	router := v1.NewRouter(app, view)
	router.InitRoutes()
	app.Get("/ping", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })
}
