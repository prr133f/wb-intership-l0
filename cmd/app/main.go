package main

import (
	"context"
	"fmt"
	"l0/internal/broker"
	"l0/internal/database"
	"l0/internal/routes"
	"l0/pkg/logger"
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logger := logger.NewZap()
	pg := database.NewPostgres(context.Background(), os.Getenv("POSTGRES_DSN"), logger)
	nats := broker.NewBroker(os.Getenv("NATS_DSN"), logger, database.NewDatabase(logger, pg))

	app := fiber.New(fiber.Config{
		AppName: "WB-INTERSHIP-L0",
	})

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	routes.InitRouter(app, logger, pg)

	logger.Fatal("NATS listener is crashed", zap.Error(nats.Listen()))
	logger.Fatal("App is crashed", zap.Error(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))))
}
