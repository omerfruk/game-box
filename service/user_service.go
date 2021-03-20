package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func CreateUser(user models.User) error {
	password := Sha256String(user.Password)
	user.Password = password
	err := database.DB().Where("mail = ? ", user.Mail).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&user).Error
		if err != nil {
			fmt.Println(err)
			fmt.Println("boyle bir kayit var")
		}
	}
	return err
}
func EditUser(user models.User) {
	database.DB().Where("id = ?", user.ID).Save(&user)
}

//kullanici score ekleme alani
func AddScore() {

}

//model ile kullaniciyi silme
func DeleteUser(user models.User) {
	if err := database.DB().Delete(&user); err != nil {
		fmt.Println(err)
	}
}
