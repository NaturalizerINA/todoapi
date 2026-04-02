package routes

import (
	"todoapi/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, noteController *controller.NoteController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/notes", noteController.Create)
	v1.Get("/notes", noteController.GetAll)
	v1.Get("/notes/:id", noteController.GetByID)
	v1.Put("/notes/:id", noteController.Update)
	v1.Delete("/notes/:id", noteController.Delete)
}
