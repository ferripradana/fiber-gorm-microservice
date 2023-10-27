package errors

import (
	"github.com/gofiber/fiber/v2"
)

type MessageResponse struct {
	Message string `json:"response"`
}

func Handler(ctx *fiber.Ctx) error {
	err := ctx.Next()

	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
		}
		ctx.Status(statusCode)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil
}
