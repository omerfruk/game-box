package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/handlers"
)

func Router(app *fiber.App) {
	//static dosyalari cekmek icin
	app.Static("/", "./")

	app.Get("/", handlers.IndexGet)
	//signup post ve get metodlari
	app.Get("/signup", handlers.SignupGet)
	app.Post("/signup", handlers.SignupPost)
	//login yonlendirmesi
	app.Get("/login", handlers.LoginGet)
	app.Post("/login", handlers.LoginPost)
	app.Get("/logout", handlers.Logout)
	//developers and developer
	app.Get("/developers", handlers.Developers)

}
