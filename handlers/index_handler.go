package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/service"
)

func IndexGet(c *fiber.Ctx) error {
	tempDevelopers := service.GetAllDevelopers()

	return c.Render("index", tempDevelopers)
}
