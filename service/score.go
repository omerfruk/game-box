package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/models"
	"gorm.io/gorm"
)

func GetScore() []models.Score {
	var temp []models.Score

	err := database.DB().Find(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err, "\n boyle bir kayit bulunamadi")
	}
	return temp
}
