package routes

import "github.com/gofiber/fiber/v2"

func ApiV1(app *fiber.App) {
	router := app.Group("/api/v1")

	router.Get("/health", Health)
	router.Get("/hello", Hello)

	app.Use("/", NotFound)
}
