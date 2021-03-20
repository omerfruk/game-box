package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/service"
)

func Developers(c *fiber.Ctx) error {
	temp := service.GetAllDevelopers()
	return c.Render("developers", temp)
}
