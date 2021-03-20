package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func CreateUser(name string, mail string, pass string)  error{
	password := Sha256String(pass)
	temp := models.User{
		Account: models.Account{
			Fullname:  name,
			Mail:      mail,
			Password:  password,
			Authority: 0,
		},
	}
	err := database.DB().Where("mail = ? ", temp.Mail).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&temp).Error
		if err != nil {
			fmt.Println(err)
			fmt.Println("boyle bir kayit var")
		}
	}
	return err
}
func AddScore() {

}
