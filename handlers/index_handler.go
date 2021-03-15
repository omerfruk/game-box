package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

func IndexGet(c *fiber.Ctx) error {
	tempDevelopers := service.GetAllDevelopers()
	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	a := models.Islogin{
		Developer: tempDevelopers,
		Bool:      !sess.Fresh(),
	}
	return c.Render("index", a)
}
