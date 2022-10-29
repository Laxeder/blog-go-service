package routes

import (
	"blog/pkg/modules/response"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Health(ctx *fiber.Ctx) error {
	now := time.Now().UTC()
	message := fmt.Sprintf("UP: %s", now.Format("2006-01-02 15:04:05"))

	return response.New(ctx).Json(response.Success(200, message))
}
