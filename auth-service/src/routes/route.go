package routes

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/user/signup")
	route.Post("/user/login")
}