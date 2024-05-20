package view

import (
	"l0/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type View struct {
	Log    *zap.Logger
	Domain domain.IFace
}

func NewView(log *zap.Logger, domain domain.IFace) IFace {
	return &View{
		Log:    log,
		Domain: domain,
	}
}

type IFace interface {
	GetDataByID(c *fiber.Ctx) error
}
