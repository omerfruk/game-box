package routers

import (
	"fmt"
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
	//developers yönlendiricisi
	app.Get("/developers", handlers.Developers)

	//session kontrol
	add := app.Group("/account", sessionControl)
	//dashboard yonlendiricisi
	add.Get("/dashboard", handlers.DashboardGet)
	//developer yönlendirici
	add.Get("/developer", handlers.Developers)
}
func sessionControl(c *fiber.Ctx) error {
	sess, err := handlers.Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	if sess.Get("name") == nil {
		return c.Redirect("/login")
	}
	return c.Next()

}
