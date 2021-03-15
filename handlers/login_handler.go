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

var Store = session.New(session.Config{
	Expiration:   5 * time.Minute,
	CookieName:   "session_id",
	KeyGenerator: utils.UUID,
})

func LoginPost(c *fiber.Ctx) error {
	var data models.User

	var temp models.Account

	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DB().Where("mail = ? and password = ?", temp.Mail, temp.Password).First(&data).Error
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/down")
	}
	Sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	defer Sess.Save()
	Sess.Set("deger", data.Fullname)
	Sess.Get("deger")
	fmt.Println(Sess)
	// session olusturulacak
	//giris kontreolu yapılacak
	//session yoketme olusturlacak
	return c.Redirect("/")
}

// logaut ediliyor
func Logout(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sess)
	//eger session acık ise girilecek yer
	// Destroy session
	if err := sess.Destroy(); err != nil {
		panic(err)
	}

	return c.Redirect("/")
}
