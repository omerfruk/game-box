package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/handlers"
)

func Router(app *fiber.App) {
	app.Get("/", handlers.IndexGet)
	//signup post ve get metodlari
	app.Get("/signup", handlers.SignupGet)
	app.Post("/signup", handlers.SignupPost)
}
