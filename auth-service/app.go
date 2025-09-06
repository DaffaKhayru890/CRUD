package main

import (
	"github.com/DaffaKhayru890/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Route(app)

	app.Listen(":3001")
}