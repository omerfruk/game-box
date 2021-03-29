package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func DeveloperGet(c *fiber.Ctx) error {
	key := c.Params("key")
	var temp models.Developer
	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DB().Where("id = ?", key).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("boyle bir kayit bulunamadi")
		return c.Redirect("/")
	}
	d := models.Islogin{
		Developer: temp,
		Bool:      !sess.Fresh(),
	}
	return c.Render("developer", d)
}
