package service

import (
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
)

func GetDeveloperAll() []models.Developer {
	var temp []models.Developer
	err := database.DB().Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}
