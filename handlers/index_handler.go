package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

func IndexGet(c *fiber.Ctx) error {
	tempDevelopers := service.GetDeveloperAll()
	sess := service.GetSesion(c)
	a := models.Islogins{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("index", a)
}
