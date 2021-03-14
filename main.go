package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/serhatkarakoca/game-box/database"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())

	database.ConnectAndMigrate()

}
