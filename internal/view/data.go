package view

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (v *View) GetDataByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := v.Domain.GetDataByIDFromCache(id)
	if err != nil {
		v.Log.Error("No such data in cache", zap.Error(err))
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(data)
}
