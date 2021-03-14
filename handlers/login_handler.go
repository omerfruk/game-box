package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
)

func LoginGet(c *fiber.Ctx) error {

	return c.Render("login", true)
}
func LoginPost(c *fiber.Ctx) error {
	var temp models.Account
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	// session olusturulacak
	//giris kontreolu yapılacak
	//session yoketme olusturlacak
	return c.Render("", temp)
}

func loginKontrol(model models.Account) {

}
