package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/game-box/database"
	"github.com/omerfruk/game-box/routers"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		MaxAge:           86400,
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, UPDATE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	database.ConnectAndMigrate()
	routers.Router(app)

	port := "4747"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}
