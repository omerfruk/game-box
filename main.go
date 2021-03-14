package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"

	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/routers"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())

	database.ConnectAndMigrate()
	routers.Router(app)

	/*
		service.CreateDeveloper("Omer Faruk TASDEMİR", "omer@mail.com", "Back-end Developer", "github.com/omerfruk", "instagram.com/omer_fruk", "https://www.linkedin.com/in/%C3%B6mer-faruk-ta%C5%9Fdemir-255114183/", "../img/developers/omer.jpg")
		service.CreateDeveloper("Serhat KARAKOCA", "serhatkarakoca@mail.com", "Game Developer", "github.com/serhatkarakoca", "instagram.com/serhat.karakoca", "https://www.linkedin.com/in/serhatkarakoca", "../img/developers/serhat.jpg")
		service.CreateDeveloper("İbrahim CAKAR", "cakaribrahim37@mail.com", "Full-Stack Developer", "github.com/ibrahimcakar", "instagram.com/ibocakar", "https://www.linkedin.com/in/ibrahim-çakar-825041183//", "../img/developers/ibrahim.jpg")
		service.CreateDeveloper("Emrullah KUTLAR", "emrullah.ktlr@mail.com", "Front-end Developer", "https://github.com/EmrullahKutlar", "instagram.com/emrllah_k", "https://www.linkedin.com/", "../img/developers/emrullah.jpg")
		service.CreateDeveloper("Osman YILMAZ", "osman.yilmaz@hotmail.com", " Junior Software Developer", "github.com/osmanyılmaz", "instagram.com/osman.yılmaz","https://www.linkedin.com/","../img/developers/osman.jpg")
		service.CreateDeveloper("musa yıldırım", "musayildirim.yldrm@gmail.com", " Junior Software Developer", "github.com/Sunset2226", "instagram.com/musayldr.m","https://www.linkedin.com/","../img/developers/musa.jpg")

	*/

	log.Fatal(app.Listen(":4747"))

}
