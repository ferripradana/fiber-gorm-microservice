package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(ctx *fiber.Ctx, req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}
