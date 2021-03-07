package service

import "github.com/omerfruk/game-box/models"

func CreateUser(name string,mail string,password string,isDeveloper bool)  {
	temp:=new(models.User)
	temp.Account.Fullname=name
	temp.Account.Mail=mail
	temp.Account.Password=password
	/*temp.Account.Userıd=*/	
}
func AddScore()  {
	
}
