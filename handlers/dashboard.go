package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/service"
)

func DashboardGet(c *fiber.Ctx) error {
	developers := service.GetDeveloperAll()
	return c.Render("dashboard", developers)
}
