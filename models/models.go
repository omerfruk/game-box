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
	gorm.Model
	ScoreID   uint
	Score     Score
	AccountID uint
	Account   Account
}

type Score struct {
	gorm.Model
	Game1 int `json:"game_1"`
	Game2 int `json:"game_2"`
	Game3 int `json:"game_3"`
}

type Developer struct {
	gorm.Model
	Duty      string `json:"duty"`
	Imgsrc    string `json:"imgsrc"`
	Inssrc    string `json:"inssrc"`
	Gitsrc    string `json:"gitsrc"`
	Linkın    string `json:"linkın"`
	AccountID uint
	Account   Account
}
