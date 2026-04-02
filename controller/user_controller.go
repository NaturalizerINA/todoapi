package controller

import (
	"todoapi/models"
	"todoapi/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req models.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request payload: " + err.Error(),
			Data:    fiber.Map{},
		})
	}

	res, err := c.service.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
			Status:  fiber.StatusUnauthorized,
			Message: err.Error(),
			Data:    fiber.Map{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:  fiber.StatusOK,
		Message: "Login successful",
		Data:    res,
	})
}
