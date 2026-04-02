package main

import (
	"log"
	"os"

	"todoapi/config"
	"todoapi/controller"
	"todoapi/repository"
	"todoapi/routes"
	"todoapi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 0. Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Proceeding with default OS environments.")
	}

	// 1. Setup Database Connection
	config.ConnectDB()

	// 2. Setup Dependency Injection (SOLID's Dependency Inversion Principle)
	noteRepo := repository.NewNoteRepository(config.DB)
	noteService := service.NewNoteService(noteRepo)
	noteController := controller.NewNoteController(noteService)

	// 3. Initialize Fiber App
	app := fiber.New()

	// 4. Setup Routes
	routes.SetupRoutes(app, noteController)

	// 5. Start the Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
