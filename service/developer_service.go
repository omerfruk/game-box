package service

import (
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
)

func CreateDeveloper(name string, mail string, duty string, gitsrc string, instsrc string, linksrc string, imgsrc string) {
	temp := new(models.Developer)
	temp.Account.Authority = 1
	temp.Imgsrc = imgsrc
	temp.Account.Fullname = name
	temp.Account.Mail = mail
	temp.Duty = duty
	temp.Gitsrc = gitsrc
	temp.Linkın = linksrc
	temp.Inssrc = instsrc
	if tempDeveloper := database.DB().Where("fullname = ?", name).Find(&temp); tempDeveloper.Error != nil {
		database.DB().Create(&temp)
	}
}
