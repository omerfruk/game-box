package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
)

func ControlLogin(temp models.Account, c *fiber.Ctx) models.User {
	var data models.User
	err := database.DB().Where("mail = ? and password = ?", temp.Mail, temp.Password).First(&data).Error
	if err != nil {
		fmt.Println(err)
		c.Redirect("/down")
	}
	return data
}
