package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func IndexGet(c *fiber.Ctx) error {

	return c.Render("index", true)
}
