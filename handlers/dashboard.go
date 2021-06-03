package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/service"
)

func Getleaderboard(c *fiber.Ctx) error {
	leaderboard := service.GetScore()
	developers := service.GetDeveloperAll()
	return c.Render("leaderboard", fiber.Map{
		"L":    leaderboard,
		"D":    developers,
		"Bool": true,
	})
}
