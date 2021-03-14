package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

<<<<<<< HEAD
func CreateUser(name string, mail string, password string, isDeveloper bool) {
	temp := new(models.User)
	temp.Account.Fullname = name
	temp.Account.Mail = mail
	temp.Account.Password = password
=======
func CreateUser(name string, mail string, password string) {
	temp := models.User{
		Account: models.Account{
			Fullname:  name,
			Mail:      mail,
			Password:  password,
			Authority: 0,
		},
	}
	err := database.DB().Where("fullname = ? ", temp.Fullname).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&temp).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("boyle bir kayit var")
	}
>>>>>>> ed6173dc6f946851a1da594d8d145cc30f29a50a
}
func AddScore() {

}
