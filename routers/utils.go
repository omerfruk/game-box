package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/game-box/context"
	"github.com/omerfruk/game-box/database"
)

func CtxWrap(h func(ctx *context.AppCtx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(&context.AppCtx{Ctx: c, Db: database.DB()})
	}
}
