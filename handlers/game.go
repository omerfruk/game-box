package handlers

import "github.com/gofiber/fiber/v2"

func SnakeGame(c *fiber.Ctx) error {
	return c.Render("snake", fiber.Map{
		"Bool": true,
	})
}
func Sayitoplama(c *fiber.Ctx) error {
	return c.Render("sayi_topla", true)
}
