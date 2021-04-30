package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

var Store = session.New(session.Config{
	Expiration:   24 * 365 * time.Hour,
	CookieName:   "session_id",
	KeyGenerator: utils.UUID,
})

func GetSesion(c *fiber.Ctx) *session.Session {

	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	return sess
}
func DestroySession(sess *session.Session) {
	err := sess.Destroy().Error()
	if err != "" {
		fmt.Println(err)
	}
}
