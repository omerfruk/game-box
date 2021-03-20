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
	fmt.Println(temp.Mail)
	if err != nil {
		fmt.Println(err)
	}
	if err =service.CreateUser(temp);err !=nil{
		fmt.Println(err)
	}
	return c.Redirect("/")
}

func SignupGet(c *fiber.Ctx) error {

	return c.Render("signup", true)
}
