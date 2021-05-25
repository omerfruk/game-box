package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func CreateDeveloper(name string, mail string, duty string, gitsrc string, instsrc string, linksrc string, imgsrc string) {

	temp := models.Developer{
		Account: models.Account{
			Fullname:  name,
			Mail:      mail,
			Password:  "",
			Authority: 1,
		},
		Duty:   duty,
		Imgsrc: imgsrc,
		Inssrc: instsrc,
		Gitsrc: gitsrc,
		Linkın: linksrc,
	}
	err := database.DB().Where("fullname = ?", temp.Fullname).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&temp).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("boyle bir kayit var")
	}
}

func GetDeveloperById(id uint) models.Developer {
	var temp models.Developer
	err := database.DB().Where("ID = ? ", id).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("boyle bi kayit yok ")
	}
	return temp
}

func UpdateDevelopers(developer models.Developer) models.Developer {
	temp := models.Developer{
		Account: models.Account{
			Fullname:  developer.Fullname,
			Mail:      developer.Mail,
			Password:  developer.Password,
			Authority: 1,
		},
		Duty:   developer.Duty,
		Imgsrc: developer.Imgsrc,
		Inssrc: developer.Inssrc,
		Gitsrc: developer.Gitsrc,
		Linkın: developer.Gitsrc,
	}
	return temp
}

func DeleteDeveloper(id string) {
	var developer models.Developer
	err := database.DB().Where("id = ?", id).Delete(&developer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("boyle bir kayit bulunamadi")
	}
}
