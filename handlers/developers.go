package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

func Developers(c *fiber.Ctx) error {
	sess := service.GetSesion(c)
	temp := service.GetDeveloperAll()
	d := models.Islogins{
		Developer: temp,
		Bool:      !sess.Fresh(),
	}
	return c.Render("developers", d)
}
