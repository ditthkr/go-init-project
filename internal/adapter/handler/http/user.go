package http

import (
	"github.com/gofiber/fiber/v2"
	"project/internal/port"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req port.UserUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user, err := r.svc.CreateUser(req)
	if err != nil {
		return handleError(c, err)
	}
	return handleSuccess(c, user)
}
