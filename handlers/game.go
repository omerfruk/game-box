package handlers

import "github.com/gofiber/fiber/v2"

func SnakeGame(c *fiber.Ctx) error {
	return c.Render("snake", true)
}
