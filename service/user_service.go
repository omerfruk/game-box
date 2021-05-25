package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func CreateUser(gelen models.User) {
	temp := models.User{
		Account: models.Account{
			Fullname:  gelen.Fullname,
			Mail:      gelen.Mail,
			Username:  gelen.Username,
			Password:  gelen.Password,
			Authority: 0,
		},
	}
	err := database.DB().Where("fullname = ? ", temp.Fullname).First(&temp).Error
	temp.Password = Sha256String(temp.Password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&temp).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("boyle bir kayit var")
	}
}

func GetUser(id string) models.User {
	var temp models.User
	err := database.DB().Where("id = ?", id).Find(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("boyle bi kullanici yok")
	}
	return temp
}

func AddScore(score models.Score, id string) {
	temp := GetUser(id)
	temp.ScoreID = &score.ID
}
func DeleteUser(id string) {
	var temp models.User
	err := database.DB().Where("id = ?", id).Delete(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

	}
}

func UpdateUser(user models.User) models.User {
	temp := models.User{
		Account: models.Account{
			Fullname:  user.Fullname,
			Mail:      user.Mail,
			Password:  user.Password,
			Authority: 0,
		},
		ScoreID: nil,
		Score:   models.Score{},
	}
	return temp
}
