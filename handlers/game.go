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
func Mario(c *fiber.Ctx) error {
	return c.Render("mario", true)
}
func BlackJack(c *fiber.Ctx) error {
	return c.Render("blackjack", true)
}
func Dino(c *fiber.Ctx) error {
	return c.Render("dino", true)
}
