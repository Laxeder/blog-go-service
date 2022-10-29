package routes

import (
	"blog/pkg/modules/response"

	"github.com/gofiber/fiber/v2"
)

func Hello(ctx *fiber.Ctx) error {
	return response.New(ctx).Json(response.Success(200, "Hello World!"))
}
