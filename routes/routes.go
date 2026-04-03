package routes

import (
	"todoapi/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App, 
	noteController *controller.NoteController,
	userController *controller.UserController,
	subtaskController *controller.SubtaskController,
) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth Routes
	v1.Post("/login", userController.Login)
	v1.Post("/register", userController.Register)

	// Note Routes
	v1.Post("/notes", noteController.Create)
	v1.Get("/notes", noteController.GetAll)
	v1.Get("/notes/:id", noteController.GetByID)
	v1.Put("/notes/:id", noteController.Update)
	v1.Delete("/notes/:id", noteController.Delete)

	// Subtask Routes
	v1.Post("/subtasks", subtaskController.Create)
	v1.Put("/subtasks/:id", subtaskController.Update)
	v1.Delete("/subtasks/:id", subtaskController.Delete)
	v1.Patch("/subtasks/:id/toggle", subtaskController.Toggle)
}
