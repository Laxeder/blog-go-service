package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	now := time.Now().UTC()
	result := fmt.Sprintf("UP: %s", now.Format("2006-01-02 15:04:05"))

	return c.Status(fiber.StatusOK).JSON(result)
}
