package controller

import (
	"strconv"
	
	"todoapi/models"
	"todoapi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type NoteController struct {
	service service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{service}
}

func (c *NoteController) Create(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	note := new(models.MasterNote)
	if err := ctx.BodyParser(note); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse request body: " + err.Error(),
			Data:    fiber.Map{},
		})
	}
	note.UserID = userID

	if err := c.service.CreateNote(note); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to create note: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:  fiber.StatusCreated,
		Message: "Note created successfully",
		Data:    note,
	})
}

func (c *NoteController) GetAll(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	notes, err := c.service.GetAllNotes(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to retrieve notes: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Success fetching all notes",
		Data:    notes,
	})
}

func (c *NoteController) GetByID(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid ID format",
			Data:    fiber.Map{},
		})
	}

	note, err := c.service.GetNoteByID(id, userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:  fiber.StatusNotFound,
			Message: "Note not found",
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    note,
	})
}

func (c *NoteController) Update(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid ID format",
			Data:    fiber.Map{},
		})
	}

	var updateData models.MasterNote
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Cannot parse request body: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	updatedNote, err := c.service.UpdateNote(id, userID, updateData)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "record not found" {
			status = fiber.StatusNotFound
		}
		return ctx.Status(status).JSON(models.APIResponse{
			Status:  status,
			Message: "Failed to update note: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Note updated successfully",
		Data:    updatedNote,
	})
}

func (c *NoteController) Delete(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid ID format",
			Data:    fiber.Map{},
		})
	}

	err = c.service.DeleteNote(id, userID)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "record not found" {
			status = fiber.StatusNotFound
		}
		return ctx.Status(status).JSON(models.APIResponse{
			Status:  status,
			Message: "Failed to delete note: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Note successfully deleted",
		Data:    fiber.Map{},
	})
}
