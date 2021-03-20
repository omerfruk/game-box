package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func CreateDeveloper(developer models.Developer) {
	err := database.DB().Where("fullname = ?", developer.Fullname).First(&developer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&developer).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("boyle bir kayit var")
	}
}
func DeleteDeveloper(developer models.Developer)  {
	database.DB().Delete(&developer)
}

func GetAllDevelopers() []models.Developer {
	var temp []models.Developer
	database.DB().Find(&temp)
	return temp
}
