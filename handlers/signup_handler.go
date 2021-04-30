package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

// post request geldigi zaman buraya yonlenecek
func SignupPost(c *fiber.Ctx) error {
	var temp models.User
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	service.CreateUser(temp)
	return c.Redirect("/")
}

func SignupGet(c *fiber.Ctx) error {
	Sess := service.GetSesion(c)
	if Sess.Fresh() != true {
		return c.Redirect("/")
	}
	return c.Render("signup", true)
}
