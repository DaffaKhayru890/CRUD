package routes

import (
	"github.com/DaffaKhayru890/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/contacts", controllers.GetContacts)
	route.Get("/contact/:id", controllers.GetContact)
	route.Post("/contact", controllers.PostContact)
	route.Patch("/contact", controllers.UpdateContacts)
	route.Delete("/contact/:id", controllers.DeleteContacts)
}