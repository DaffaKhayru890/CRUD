package main

import (
	"log"

	"github.com/DaffaKhayru890/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.Route(app)

	app.Listen(":3000")
}