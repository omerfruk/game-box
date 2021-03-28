package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func Developers(c *fiber.Ctx) error {
	var temp []models.Developer
	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DB().Find(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
	}
	d := models.Islogin{
		Developer: temp,
		Bool:      !sess.Fresh(),
	}
	return c.Render("developers", d)
}
