package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (middleware *Server) fetchUserId(ctx *fiber.Ctx) error {
	header := ctx.Request().Header.Peek("X-User-Id")

	id, err := strconv.ParseUint(string(header), 10, 64)
	if err != nil {

	}

	if id == 0 {
	}

	ctx.Locals("id", id)

	return ctx.Next()
}
