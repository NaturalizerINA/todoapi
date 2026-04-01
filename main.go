package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MasterNote represents the existing master_notes table
type MasterNote struct {
	ID     int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Status string `gorm:"column:status" json:"status"`
}

// TableName tells GORM to use the exactly named table `master_notes`
func (MasterNote) TableName() string {
	return "master_notes"
}

// APIResponse represents the standardized JSON response structure
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Global variable for DB connection
var DB *gorm.DB

func main() {
	// Initialize Database Connection
	// Default postgres port is 5432 and user is 'postgres'
	dsn := "host=localhost user=postgres password=M1r4etft22! dbname=notes port=5432 sslmode=disable TimeZone=UTC"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection successful!")

	// Initialize Fiber app
	app := fiber.New()

	// Setup Routes
	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3009"))
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// CRUD for target table: master_notes
	v1.Post("/notes", CreateNote)
	v1.Get("/notes", GetNotes)
	v1.Get("/notes/:id", GetNote)
	v1.Put("/notes/:id", UpdateNote)
	v1.Delete("/notes/:id", DeleteNote)
}

// --- CRUD Handlers ---

// CreateNote creates a new entry in master_notes
func CreateNote(c *fiber.Ctx) error {
	note := new(MasterNote)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse request body: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	if result := DB.Create(&note); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to create note: " + result.Error.Error(),
			Data:    fiber.Map{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Status:  fiber.StatusCreated,
		Message: "Note created successfully",
		Data:    note,
	})
}

// GetNotes retrieves all entries from master_notes
func GetNotes(c *fiber.Ctx) error {
	var notes []MasterNote
	
	if result := DB.Find(&notes); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to retrieve notes: " + result.Error.Error(),
			Data:    fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  fiber.StatusOK,
		Message: "Success fetching all notes",
		Data:    notes,
	})
}

// GetNote retrieves a single note by ID
func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note MasterNote
	
	if result := DB.First(&note, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Status:  fiber.StatusNotFound,
			Message: "Note not found",
			Data:    fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    note,
	})
}

// UpdateNote updates an existing note by ID
func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note MasterNote

	// Find the existing note
	if result := DB.First(&note, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Status:  fiber.StatusNotFound,
			Message: "Note not found",
			Data:    fiber.Map{},
		})
	}

	// Parse incoming data
	updateData := new(MasterNote)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse request body: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	// Update necessary fields
	note.Name = updateData.Name
	note.Status = updateData.Status

	if result := DB.Save(&note); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to update note: " + result.Error.Error(),
			Data:    fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  fiber.StatusOK,
		Message: "Note updated successfully",
		Data:    note,
	})
}

// DeleteNote deletes an existing note by ID
func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note MasterNote

	// Verify note exists first
	if result := DB.First(&note, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Status:  fiber.StatusNotFound,
			Message: "Note not found",
			Data:    fiber.Map{},
		})
	}

	if result := DB.Delete(&note); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to delete note: " + result.Error.Error(),
			Data:    fiber.Map{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  fiber.StatusOK,
		Message: "Note successfully deleted",
		Data:    fiber.Map{},
	})
}
