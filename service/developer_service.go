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

func GetAllDevelopers() []models.Developer {
	var temp []models.Developer
	database.DB().Find(&temp)
	return temp
}
