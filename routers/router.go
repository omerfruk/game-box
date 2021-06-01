package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/handlers"
	"github.com/omerfruk/game-box/service"
)

func Router(app *fiber.App) {
	//static dosyalari cekmek icin
	app.Static("/", "./")

	app.Get("/", handlers.IndexGet)

	//signup post ve get metodlari yonlendirmesi
	app.Get("/signup", handlers.SignupGet)
	app.Post("/signup", handlers.SignupPost)

	//login sayfasi yonlendirmesi
	app.Get("/login", handlers.LoginGet)
	app.Post("/login", handlers.LoginPost)
	app.Get("/logout", handlers.Logout)

	//oyun sayfasina yönlendirme
	game := app.Group("/game")
	game.Get("/snake", handlers.SnakeGame)
	game.Get("/sayi-toplama", handlers.Sayitoplama)
	game.Get("/mario", handlers.Mario)
	game.Get("/black-jack", handlers.BlackJack)
	game.Get("/dino", handlers.Dino)

	//developers sayfasina yönlendiricisi
	app.Get("/developers", handlers.Developers)
	app.Get("/developer/:key", handlers.DeveloperGet)
	app.Get("/account/:key", handlers.GetAccount)

	//session kontrol
	add := app.Group("/account", sessionControl)

	//dashboard yonlendiricisi
	add.Get("/dashboard", handlers.DashboardGet)
	//developer yönlendirici
	//add.Get("/developer/:key", handlers.DeveloperGet)
}
func sessionControl(c *fiber.Ctx) error {
	sess := service.GetSesion(c)
	if sess.Get("name") == nil {
		return c.Redirect("/login")
	}
	return c.Next()

}
