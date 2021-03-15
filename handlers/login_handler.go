package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"time"
)

func LoginGet(c *fiber.Ctx) error {

	return c.Render("login", true)
}

var store = session.New(session.Config{
	Expiration:   5 * time.Minute,
	CookieName:   "session_id",
	KeyGenerator: utils.UUID,
})

func LoginPost(c *fiber.Ctx) error {
	var deger models.User

	var temp models.Account

	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DB().Where("mail = ? and password = ?", temp.Mail, temp.Password).First(&deger).Error
	if err != nil {
		fmt.Println(err)
	}
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	defer sess.Save()
	sess.Set("deger", deger.Fullname)
	sess.Get("deger")

	// session olusturulacak
	//giris kontreolu yapılacak
	//session yoketme olusturlacak
	return c.Redirect("/")
}
