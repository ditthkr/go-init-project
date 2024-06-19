package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"project/internal/adapter/errs"
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
		var isErr errs.Err
		ok := errors.As(err, &isErr)
		if ok {
			return c.Status(isErr.Code).JSON(fiber.Map{"message": isErr.Message})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
