package database

import (
	"fmt"
	"github.com/omerfruk/game-box/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	HOST     = "pencereler.hayrat.dev"
	DATABASE = "teknolojilideri.xyz"
	USER     = "teknolojilideri.xyz"
	PASSWORD = "teknolojilideri.xyz1442?"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

//db baglantısını gerceklestirecegimiz alan
func Connect() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	var err error
	db, err = gorm.Open(postgres.Open(vt), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
}

//2 methodu birden cagirma metodu
func ConnectAndMigrate() {
	Connect()
	Migrate()
}

//db ile models verilerini haberdar edecek method
func Migrate() {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Account{})
	db.AutoMigrate(models.Developer{})
	db.AutoMigrate(models.Score{})
}
