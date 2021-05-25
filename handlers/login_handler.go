package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/models"
	"github.com/omerfruk/game-box/service"
)

func LoginGet(c *fiber.Ctx) error {
	Sess := service.GetSesion(c)
	if Sess.Fresh() != true {
		return c.Redirect("/")
	}
	return c.Render("login", true)
}

func LoginPost(c *fiber.Ctx) error {
	// Body parse etmek icin model olusturalım
	var temp models.Account
	// formdan gelen model burda parse ediliyor
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	//gellen verilerin kontrolu ve hash edildi
	temp.Password = service.Sha256String(temp.Password)
	data := service.ControlLogin(temp, c)
	Sess := service.GetSesion(c)
	defer Sess.Save()
	Sess.Set("name", data.Fullname)
	Sess.Set("idsi", data.ID)
	//veriyi geri dondur
	return c.Redirect("/")
}

// logaut ediliyor
func Logout(c *fiber.Ctx) error {
	//session olusturuldu
	sess := service.GetSesion(c)
	// Destroy session
	service.DestroySession(sess)
	return c.Redirect("/")
}
