package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

func SnakeGame(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("snake", a)
}
func Sayitoplama(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("sayi_topla", a)
}
func Mario(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("mario", a)
}
func BlackJack(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("blackjack", a)
}
func Dino(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("dino", a)
}
