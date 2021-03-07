package models

import (
	"gorm.io/gorm"
)

type Authority int

const (
	Gamer Authority = iota + 1
	developer
)

type Account struct {
	gorm.Model
	Fullname  string    `json:"fullname"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	Authority Authority `json:"authority"`
}

type User struct {
	AccountId uint
	Account Account
	Score  []Score
}

type Score struct {
	Game1   int `json:"game_1"`
	Game2   int `json:"game_2"`
	Game3   int `json:"game_3"`
	GamerID int `json:"gamer_id"`
}

type Developer struct {
	gorm.Model
	AccountId uint
	Duty string `json:"duty"`
	Imgsrc string `json:"imgsrc"`
	Gitsrc string `json:"gitsrc"`
	Linkın string `json:"linkın"`
	Account Account
}
