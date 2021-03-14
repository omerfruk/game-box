package service

import "github.com/omerfruk/game-box/models"

func CreateUser(name string, mail string, password string, isDeveloper bool) {
	temp := new(models.User)
	temp.Account.Fullname = name
	temp.Account.Mail = mail
	temp.Account.Password = password
	if isDeveloper == true {
		temp.Account.Authority = 1
	} else {
		temp.Account.Authority = 2
	}
}
func AddScore() {

}
