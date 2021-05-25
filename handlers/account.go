package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/service"
	"strconv"
)

func GetAccount(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("key"), 10, 32)
	account := service.GetDeveloperById(uint(id))
	return c.Render("account", fiber.Map{
		"Account": account,
		"Bool":    true,
	})
}
