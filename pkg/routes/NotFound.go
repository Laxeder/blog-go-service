package routes

import (
	"blog/pkg/modules/response"

	"github.com/gofiber/fiber/v2"
)

func NotFound(ctx *fiber.Ctx) error {
	return response.New(ctx).Json(response.Error(404, "BGS001", "Recurso não encontrado."))
}
