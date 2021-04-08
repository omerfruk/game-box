package models

import (
	"gorm.io/gorm"
)

type Authority int

const (
	Gamer Authority = iota + 1
	developer
)

type Islogins struct {
	Developer []Developer
	Bool      bool
}

//tek developer icin gecerli
type Islogin struct {
	Developer Developer
	Bool      bool
}
type Account struct {
	gorm.Model
	Fullname  string    `json:"fullname"`
	Mail      string    `json:"mail"`
	Username  string    `json:"user_name"`
	Password  string    `json:"password"`
	Authority Authority `json:"authority"`
}

type User struct {
	Account
	ScoreID *uint
	Score   Score
}

type Score struct {
	gorm.Model
	Game1 int `json:"game_1"`
	Game2 int `json:"game_2"`
	Game3 int `json:"game_3"`
}

type Developer struct {
	Account
	Duty   string `json:"duty"`
	Imgsrc string `json:"imgsrc"`
	Inssrc string `json:"inssrc"`
	Gitsrc string `json:"gitsrc"`
	Linkın string `json:"linkın"`
}
