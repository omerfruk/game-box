package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
	"strconv"
)

func DeveloperGet(c *fiber.Ctx) error {
	key, _ := strconv.ParseUint(c.Params("key"), 10, 32)
	sess := service.GetSesion(c)
	temp := service.GetDeveloperById(uint(key))
	d := models.Islogin{
		Developer: temp,
		Bool:      !sess.Fresh(),
	}
	return c.Render("account", d)
}
