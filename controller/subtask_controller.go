package controller

import (
	"strconv"
	"todoapi/models"
	"todoapi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SubtaskController struct {
	service service.SubtaskService
}

func NewSubtaskController(service service.SubtaskService) *SubtaskController {
	return &SubtaskController{service}
}

func (c *SubtaskController) Create(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	var subtask models.Subtask
	if err := ctx.BodyParser(&subtask); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid input",
			Data:    nil,
		})
	}

	if err := c.service.Create(&subtask, userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:  fiber.StatusCreated,
		Message: "Subtask created",
		Data:    subtask,
	})
}

func (c *SubtaskController) Update(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, _ := strconv.Atoi(ctx.Params("id"))
	var subtask models.Subtask
	if err := ctx.BodyParser(&subtask); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid input",
			Data:    nil,
		})
	}
	subtask.ID = id
	if err := c.service.Update(&subtask, userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Subtask updated",
		Data:    subtask,
	})
}

func (c *SubtaskController) Delete(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := c.service.Delete(id, userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Subtask deleted",
		Data:    nil,
	})
}

func (c *SubtaskController) Toggle(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := c.service.Toggle(id, userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Subtask toggled",
		Data:    nil,
	})
}
