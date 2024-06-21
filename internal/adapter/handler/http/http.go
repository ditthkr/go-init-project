package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"project/internal/adapter/errs"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func handleError(c *fiber.Ctx, err error) error {
	var isErr errs.Err
	ok := errors.As(err, &isErr)
	if ok {
		return c.Status(isErr.Code).JSON(response{
			Success: false,
			Message: isErr.Error(),
		})

	}
	return c.Status(fiber.StatusBadRequest).JSON(response{
		Success: false,
		Message: err.Error(),
	})
}

func handleSuccess(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(response{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}
