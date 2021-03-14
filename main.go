package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/routers"
	"github.com/omerfruk/game-box/service"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())

	database.ConnectAndMigrate()
	routers.Router(app)
	service.CreateDeveloper("Ömer Faruk Taşdemir", "omer@mail.com", "Back-and Developer", "github.com/omerfruk", "instagram.com/omer_fruk", "https://www.linkedin.com/in/%C3%B6mer-faruk-ta%C5%9Fdemir-255114183/", "../img/developers/omer.jpg")
	log.Fatal(app.Listen(":4747"))

}
