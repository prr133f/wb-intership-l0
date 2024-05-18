package view

import "github.com/gofiber/fiber/v2"

func (v *View) GetDataById(c *fiber.Ctx) error {
	return c.SendString("dataById")
}
